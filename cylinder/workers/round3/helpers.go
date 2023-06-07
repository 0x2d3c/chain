package round3

import (
	"github.com/bandprotocol/chain/v2/cylinder/client"
	"github.com/bandprotocol/chain/v2/cylinder/store"
	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/x/tss/types"
)

// getOwnPrivKey calculates the own private key for the group member and verifies the secret shares.
// It returns the own private key, a slice of complaints (if any), and an error, if any.
func getOwnPrivKey(group store.Group, groupRes *client.GroupResponse) (tss.PrivateKey, []types.Complain, error) {
	round1DataI, err := groupRes.GetRound1Data(group.MemberID)
	if err != nil {
		return nil, nil, err
	}

	var secretShares tss.Scalars
	var complains []types.Complain
	for j := uint64(1); j <= groupRes.Group.Size_; j++ {
		// Calculate your own secret value
		if j == uint64(group.MemberID) {
			secretShare, err := tss.ComputeSecretShare(group.Coefficients, group.MemberID)
			if err != nil {
				return nil, nil, err
			}
			secretShares = append(secretShares, secretShare)
			continue
		}

		// Get Round1Commitment of J
		commitmentJ, err := groupRes.GetRound1Data(tss.MemberID(j))
		if err != nil {
			return nil, nil, err
		}

		// Get secret share
		secretShare, err := getSecretShare(
			group.MemberID,
			tss.MemberID(j),
			group.OneTimePrivKey,
			commitmentJ.OneTimePubKey,
			groupRes,
		)
		if err != nil {
			return nil, nil, err
		}

		// Verify secret share
		err = tss.VerifySecretShare(group.MemberID, secretShare, commitmentJ.CoefficientsCommit)
		if err != nil {
			// Generate complaint if we fail to verify secret share
			sig, keySym, nonceSym, err := tss.SignComplain(
				round1DataI.OneTimePubKey,
				commitmentJ.OneTimePubKey,
				group.OneTimePrivKey,
			)
			if err != nil {
				return nil, nil, err
			}

			// Add complaint
			complains = append(complains, types.Complain{
				I:         group.MemberID,
				J:         tss.MemberID(j),
				KeySym:    keySym,
				Signature: sig,
				NonceSym:  nonceSym,
			})

			continue
		}

		// Add secret share if verification is successful
		secretShares = append(secretShares, secretShare)
	}

	if len(complains) > 0 {
		return nil, complains, nil
	}

	ownPrivKey, err := tss.ComputeOwnPrivateKey(secretShares...)
	if err != nil {
		return nil, nil, err
	}
	return ownPrivKey, nil, nil
}

// getSecretShare calculates and retrieves the decrypted secret share between two members.
// It takes the member IDs, private and public keys, and the group response as input.
// It returns the decrypted secret share and any error encountered during the process.
func getSecretShare(
	i, j tss.MemberID,
	privKeyI tss.PrivateKey,
	pubKeyJ tss.PublicKey,
	groupRes *client.GroupResponse,
) (tss.Scalar, error) {
	// Calculate keySym
	keySym, err := tss.ComputeKeySym(privKeyI, pubKeyJ)
	if err != nil {
		return nil, err
	}

	// Get encrypted secret share between I and J
	esc, err := groupRes.GetEncryptedSecretShare(j, i)
	if err != nil {
		return nil, err
	}

	// Decrypt secret share
	secretShare, err := tss.Decrypt(esc, keySym)
	if err != nil {
		return nil, err
	}

	return secretShare, nil
}
