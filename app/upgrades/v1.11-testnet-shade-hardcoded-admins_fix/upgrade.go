package v1_11_testnet_shade_hardcoded_admins_fix

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/scrtlabs/SecretNetwork/app/keepers"
	"github.com/scrtlabs/SecretNetwork/app/upgrades"
	"github.com/scrtlabs/SecretNetwork/x/compute"
	computetypes "github.com/scrtlabs/SecretNetwork/x/compute/internal/types"
)

const upgradeName = "v1.11-testnet-shade-hardcoded-admins-fix"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          upgradeName,
	CreateUpgradeHandler: createUpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}

var hardcodedContractAdmins = map[string]string{
	"secret14svk0x3sztxwta9kv9dv6fwzqlc26mmjfyypc2": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret13yzengut04fpk0f9hs4axvyz4np30qczt0pa7z": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1e0k5jza9jqctc5dt7mltnxmwpu3a3kqe0a6hf3": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret17aku787dnktxtagrx2vp9xp2ym4wa7ktqv5h6r": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret13w6n5u3kpvqdunkavgfy40d7ma85xuhxrcxd0a": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1aeu5lcj8dhhaae406y7g4afy5wtcgvcwdpuh6n": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1x8gs5yja6f2mmvmf5thr4r7w6kp594lrhgxclt": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1umunptajd6j3j02wchdftqkhns48ysp0tguaad": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret16xak8matccjjn4k45em9fv4j28zu2c4hdw96hg": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1khkf49xfgjtqyprd39jlyqj90axyl8kw4nlmcz": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1qa0l6tt9drkf9jk9rty3f37p23ch6vpzvgetlu": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1z6l0pg9gynzgk7qsqdaj8d9nkx6w0hctukfx4v": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1qn07k2d7hcmy8kuk7d28f5evzwygwvwvqeqzhz": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1apqgmdm7d2emufxkdujwuglrgzhsskxj8xpjls": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1t5sq6mlggs04u4ukfqyhqa00h8aehf4e62f6xm": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret16dyx6yukjg6fvdwz9935glesqvw2mtujuplq9y": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret126tc6kwgwj33vqnllytjhjlrghnrrqd2llqr9y": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1lxkltxpft6suhf63x6dvyeghqlwqldz8t2wesz": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1j22chejflprk06wv9cgz9la3tm3fkjyd92s94r": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret12k8jf45n50exzu0299lalxzr3wy02yzrrxxd38": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1uler557j3xdkqu9ua637gu2lce5557grlnw0u0": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1pdk4wj2mtkpger96lky9ptjk6zmqv7f0cz256q": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1wae7v026p9q7vapatgxwdmrmpe020wlsesxkmd": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1d5qw64q68yz2qj3qgnr5f20kemyrtnghsgngql": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1h5fmyt9424cgae4jcnre70p9s05dmmyqx66lp0": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1q9le3t099ad6nh6tm0k2lqnsq59zpa9hdzwl4a": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret12vhfcdc90tygd499ecdhsg6dwfp0p6ncrl5x2d": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1d9pj42dfgnx45uwuxlup55k2fle7d0e5u94xvg": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret16f3j4kvecpeepg7cvrdu7fmj8fmpfjt52vfjh2": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret182js23ceyjywkvnxpqd6sge6v5062uh0q4gu3c": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret16076kg6k2dvypcdx4gfnmd8swquqyv23t6jz8s": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret1emph38m50343r9tj8l79quw5kpdapeaku8yzpq": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret18mf0gg96da5jjsjfaudsuh5kgmmfmjfg4r8zjj": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret15z660976c54e8apx6q83at74ekvp787qsrast8": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
	"secret18tnfyh4xfetqjdmy6f4hpzkqgtnt7vlvam7kj7": "secret1fz7ugsneqqru9pex7h4pjyn77s4fsmcp7sycyl",
}

func v1GetContractKey(ctx sdk.Context, k *compute.Keeper, contractAddress sdk.AccAddress) []byte {
	store := ctx.KVStore(k.GetStoreKey())

	contractKey := store.Get(computetypes.GetContractEnclaveKey(contractAddress))

	return contractKey
}

func setHardCodedAdmins(ctx sdk.Context, keeper *compute.Keeper) error {
	iter := prefix.NewStore(ctx.KVStore(keeper.GetStoreKey()), computetypes.ContractKeyPrefix).Iterator(nil, nil)
	for ; iter.Valid(); iter.Next() {
		var contractAddress sdk.AccAddress = iter.Key()

		var contractInfo computetypes.ContractInfo
		keeper.GetCdc().MustUnmarshal(iter.Value(), &contractInfo)

		if hardcodedContractAdmins[contractAddress.String()] != "" {
			contractInfo.Admin = hardcodedContractAdmins[contractAddress.String()]
			// When the hardcodedContractAdmins has a hardcoded admin via gov, adminProof is ignored inside the enclave.
			// Otherwise and if valid, adminProof is a 32 bytes array (output of sha256).
			// For future proofing and avoiding passing null pointers to the enclave, we'll set it to a 32 bytes array of 0.
			contractInfo.AdminProof = make([]byte, 32)
		}

		// get v1 contract key
		v1ContractKey := v1GetContractKey(ctx, keeper, contractAddress)

		// convert v1 contract key to v2 contract key
		v2ContractKey := computetypes.ContractKey{
			OgContractKey:           v1ContractKey,
			CurrentContractKey:      v1ContractKey,
			CurrentContractKeyProof: nil,
		}

		// overide v1 contract key with v2 contract key in the store
		keeper.SetContractKey(ctx, contractAddress, &v2ContractKey)
	}

	return nil
}

func createUpgradeHandler(mm *module.Manager, k *keepers.SecretAppKeepers, configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info(` _    _ _____   _____ _____            _____  ______ `)
		ctx.Logger().Info(`| |  | |  __ \ / ____|  __ \     /\   |  __ \|  ____|`)
		ctx.Logger().Info(`| |  | | |__) | |  __| |__) |   /  \  | |  | | |__   `)
		ctx.Logger().Info(`| |  | |  ___/| | |_ |  _  /   / /\ \ | |  | |  __|  `)
		ctx.Logger().Info(`| |__| | |    | |__| | | \ \  / ____ \| |__| | |____ `)
		ctx.Logger().Info(` \____/|_|     \_____|_|  \_\/_/    \_\_____/|______|`)

		setHardCodedAdmins(ctx, k.ComputeKeeper)

		ctx.Logger().Info(fmt.Sprintf("Running module migrations for %s...", upgradeName))
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
