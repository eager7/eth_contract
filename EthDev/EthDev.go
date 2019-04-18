// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_dev

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthDevABI is the input ABI used to generate the binding from.
const EthDevABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"m_numOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetSpentToday\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"m_required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_h\",\"type\":\"bytes32\"}],\"name\":\"confirm\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newLimit\",\"type\":\"uint256\"}],\"name\":\"setDailyLimit\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"_r\",\"type\":\"bytes32\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operation\",\"type\":\"bytes32\"}],\"name\":\"revoke\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRequired\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operation\",\"type\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"hasConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"kill\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"m_dailyLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"},{\"name\":\"_daylimit\",\"type\":\"uint256\"}],\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operation\",\"type\":\"bytes32\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operation\",\"type\":\"bytes32\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"}],\"name\":\"OwnerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newRequirement\",\"type\":\"uint256\"}],\"name\":\"RequirementChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SingleTransact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operation\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MultiTransact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operation\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ConfirmationNeeded\",\"type\":\"event\"}]"

// EthDevBin is the compiled bytecode used for deploying new contracts.
const EthDevBin = `0x606060405236156100b95760e060020a6000350463173825d9811461010b5780632f54bf6e1461015b5780634123cb6b146101835780635c52c2f51461018c5780637065cb48146101b2578063746c9171146101db578063797af627146101e4578063b20d30a9146101f7578063b61d27f614610220578063b75c7dc614610241578063ba51a6df14610270578063c2cf732614610299578063cbf0b0c0146102d7578063f00d4b5d14610300578063f1736d861461032e575b61033860003411156101095760408051600160a060020a033316815234602082015281517fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c929181900390910190a15b565b6103386004356000600036604051808383808284375050509081018190039020905061063d815b600160a060020a03331660009081526101026020526040812054818082811415610bb357610d0b565b61033a6004355b600160a060020a03811660009081526101026020526040812054115b919050565b61033a60015481565b610338600036604051808383808284375050509081018190039020905061078e81610132565b61033860043560003660405180838380828437505050908101819003902090506105b581610132565b61033a60005481565b61033a6004355b6000816109f181610132565b610338600435600036604051808383808284375050509081018190039020905061078281610132565b61033a6004803590602480359160443591820191013560006107ad33610162565b610338600435600160a060020a0333166000908152610102602052604081205490808281141561034c576103cb565b61033860043560003660405180838380828437505050908101819003902090506106fc81610132565b61033a600435602435600082815261010360209081526040808320600160a060020a0385168452610102909252822054828181141561075557610779565b610338600435600036604051808383808284375050509081018190039020905061079c81610132565b6103386004356024356000600036604051808383808284375050509081018190039020905061045681610132565b61033a6101055481565b005b60408051918252519081900360200190f35b50506000828152610103602052604081206001810154600284900a9290831611156103cb5780546001828101805492909101835590839003905560408051600160a060020a03331681526020810186905281517fc7fb647e59b18047309aa15aad418e5d7ca96d173ad704f1031a2c3d7591734b929181900390910190a15b50505050565b600160a060020a03831660028361010081101561000257508301819055600160a060020a03851660008181526101026020908152604080832083905584835291829020869055815192835282019290925281517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c929181900390910190a1505b505050565b156103cb5761046483610162565b1561046f5750610451565b600160a060020a0384166000908152610102602052604081205492508214156104985750610451565b6103d15b6101045460005b81811015610e5857610104805461010891600091849081101561000257600080516020610f1383398151915201548252506020918252604081208054600160a060020a0319168155600181018290556002810180548382559083528383209193610edd92601f92909201048101906109d9565b60018054810190819055600160a060020a038316906002906101008110156100025790900160005081905550600160005054610102600050600084600160a060020a03168152602001908152602001600020600050819055507f994a936646fe87ffe4f1e469d3d6aa417d6b855598397f323de5b449f765f0c3826040518082600160a060020a0316815260200191505060405180910390a15b505b50565b156105b0576105c382610162565b156105ce57506105b2565b6105d661049c565b60015460fa90106105eb576105e9610600565b505b60015460fa901061051657506105b2565b6106ba5b600060015b6001548110156109ed575b600154811080156106305750600281610100811015610002570154600014155b15610d1357600101610610565b1561045157600160a060020a03831660009081526101026020526040812054925082141561066b57506105b0565b600160016000505403600060005054111561068657506105b0565b600060028361010081101561000257508301819055600160a060020a038416815261010260205260408120556105fc61049c565b5060408051600160a060020a038516815290517f58619076adf5bb0943d100ef88d52d7c3fd691b19d3a9071b555b651fbf418da9181900360200190a1505050565b156105b05760015482111561071157506105b2565b600082905561071e61049c565b6040805183815290517facbdb084c721332ac59f9b8e392196c9eb0e4932862da8eb9beaf0dad4f550da9181900360200190a15050565b506001820154600282900a908116600014156107745760009350610779565b600193505b50505092915050565b156105b0575061010555565b156105b25760006101065550565b156105b05781600160a060020a0316ff5b156109c9576107c1846000610ded33610162565b1561087d577f92ca3a80853e6663fa31fa10b99225f18d4902939b4c53a9caae9043f6efd00433858786866040518086600160a060020a0316815260200185815260200184600160a060020a031681526020018060200182810382528484828181526020019250808284378201915050965050505050505060405180910390a184600160a060020a03168484846040518083838082843750505090810191506000908083038185876185025a03f150600093506109c992505050565b600036436040518084848082843750505090910190815260405190819003602001902091506108ad9050816101eb565b1580156108d0575060008181526101086020526040812054600160a060020a0316145b156109c95760008181526101086020908152604082208054600160a060020a03191688178155600181018790556002018054858255818452928290209092601f019190910481019084908682156109d1579182015b828111156109d1578235826000505591602001919060010190610925565b50507f1733cbb53659d713b79580f79f3f9ff215f78a7c7aa45890f3b89fc5cddfbf328133868887876040518087815260200186600160a060020a0316815260200185815260200184600160a060020a03168152602001806020018281038252848482818152602001925080828437820191505097505050505050505060405180910390a15b949350505050565b506109439291505b808211156109ed57600081556001016109d9565b5090565b15610ba05760008381526101086020526040812054600160a060020a031614610ba05760408051600091909120805460018201546002929092018054600160a060020a0392909216939091819083908015610a7157820191906000526020600020905b815481529060010190602001808311610a5457829003601f168201915b505091505060006040518083038185876185025a03f150505060008481526101086020908152604080519281902080546001820154600160a060020a033381811688529587018b905293860181905292166060850181905260a06080860181815260029390930180549187018290527fe7c957c06e9a662c1a6c77366179f5b702b97651dc28eee7d5bf1dff6e40bb4a975094958a959293909160c083019084908015610b4357820191906000526020600020905b815481529060010190602001808311610b2657829003601f168201915b5050965050505050505060405180910390a160008381526101086020908152604082208054600160a060020a031916815560018101839055600281018054848255908452828420919392610ba692601f92909201048101906109d9565b50919050565b505050600191505061017e565b60008581526101036020526040812080549093501415610c3b576000805483556001838101919091556101048054918201808255828015829011610c0a57818360005260206000209182019101610c0a91906109d9565b50505060028301819055610104805487929081101561000257600091909152600080516020610f1383398151915201555b506001810154600283900a90811660001415610d0b5760408051600160a060020a03331681526020810187905281517fe1c52dc63b719ade82e8bea94cc41a0d5d28e4aaf536adb5e9cccc9ff8c1aeda929181900390910190a1815460019011610cf8576000858152610103602052604090206002015461010480549091908110156100025760406000908120600080516020610f138339815191529290920181905580825560018281018290556002909201559450610d0b9050565b8154600019018255600182018054821790555b505050919050565b5b60018054118015610d3657506001546002906101008110156100025701546000145b15610d4a5760018054600019019055610d14565b60015481108015610d6d5750600154600290610100811015610002570154600014155b8015610d8757506002816101008110156100025701546000145b15610de857600154600290610100811015610002578101549082610100811015610002578101919091558190610102906000908361010081101561000257810154825260209290925260408120929092556001546101008110156100025701555b610605565b1561017e5761010754610e035b62015180420490565b1115610e1c57600061010655610e17610dfa565b610107555b6101065480830110801590610e3a5750610106546101055490830111155b15610e505750610106805482019055600161017e565b50600061017e565b6105b06101045460005b81811015610ee85761010480548290811015610002576000918252600080516020610f13833981519152015414610ed557610104805461010391600091849081101561000257600080516020610f1383398151915201548252506020919091526040812081815560018101829055600201555b600101610e62565b5050506001016104a3565b610104805460008083559190915261045190600080516020610f13833981519152908101906109d956004c0be60200faa20559308cb7b5a1bb3255c16cb1cab91f525b5ae7a03d02fabe`

// DeployEthDev deploys a new Ethereum contract, binding an instance of EthDev to it.
func DeployEthDev(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _required *big.Int, _daylimit *big.Int) (common.Address, *types.Transaction, *EthDev, error) {
	parsed, err := abi.JSON(strings.NewReader(EthDevABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthDevBin), backend, _owners, _required, _daylimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthDev{EthDevCaller: EthDevCaller{contract: contract}, EthDevTransactor: EthDevTransactor{contract: contract}, EthDevFilterer: EthDevFilterer{contract: contract}}, nil
}

// EthDev is an auto generated Go binding around an Ethereum contract.
type EthDev struct {
	EthDevCaller     // Read-only binding to the contract
	EthDevTransactor // Write-only binding to the contract
	EthDevFilterer   // Log filterer for contract events
}

// EthDevCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthDevCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDevTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthDevTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDevFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthDevFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDevSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthDevSession struct {
	Contract     *EthDev           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthDevCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthDevCallerSession struct {
	Contract *EthDevCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthDevTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthDevTransactorSession struct {
	Contract     *EthDevTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthDevRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthDevRaw struct {
	Contract *EthDev // Generic contract binding to access the raw methods on
}

// EthDevCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthDevCallerRaw struct {
	Contract *EthDevCaller // Generic read-only contract binding to access the raw methods on
}

// EthDevTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthDevTransactorRaw struct {
	Contract *EthDevTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthDev creates a new instance of EthDev, bound to a specific deployed contract.
func NewEthDev(address common.Address, backend bind.ContractBackend) (*EthDev, error) {
	contract, err := bindEthDev(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthDev{EthDevCaller: EthDevCaller{contract: contract}, EthDevTransactor: EthDevTransactor{contract: contract}, EthDevFilterer: EthDevFilterer{contract: contract}}, nil
}

// NewEthDevCaller creates a new read-only instance of EthDev, bound to a specific deployed contract.
func NewEthDevCaller(address common.Address, caller bind.ContractCaller) (*EthDevCaller, error) {
	contract, err := bindEthDev(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthDevCaller{contract: contract}, nil
}

// NewEthDevTransactor creates a new write-only instance of EthDev, bound to a specific deployed contract.
func NewEthDevTransactor(address common.Address, transactor bind.ContractTransactor) (*EthDevTransactor, error) {
	contract, err := bindEthDev(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthDevTransactor{contract: contract}, nil
}

// NewEthDevFilterer creates a new log filterer instance of EthDev, bound to a specific deployed contract.
func NewEthDevFilterer(address common.Address, filterer bind.ContractFilterer) (*EthDevFilterer, error) {
	contract, err := bindEthDev(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthDevFilterer{contract: contract}, nil
}

// bindEthDev binds a generic wrapper to an already deployed contract.
func bindEthDev(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthDevABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthDev *EthDevRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthDev.Contract.EthDevCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthDev *EthDevRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthDev.Contract.EthDevTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthDev *EthDevRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthDev.Contract.EthDevTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthDev *EthDevCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthDev.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthDev *EthDevTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthDev.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthDev *EthDevTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthDev.Contract.contract.Transact(opts, method, params...)
}

// HasConfirmed is a free data retrieval call binding the contract method 0xc2cf7326.
//
// Solidity: function hasConfirmed(bytes32 _operation, address _owner) constant returns(bool)
func (_EthDev *EthDevCaller) HasConfirmed(opts *bind.CallOpts, _operation [32]byte, _owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthDev.contract.Call(opts, out, "hasConfirmed", _operation, _owner)
	return *ret0, err
}

// HasConfirmed is a free data retrieval call binding the contract method 0xc2cf7326.
//
// Solidity: function hasConfirmed(bytes32 _operation, address _owner) constant returns(bool)
func (_EthDev *EthDevSession) HasConfirmed(_operation [32]byte, _owner common.Address) (bool, error) {
	return _EthDev.Contract.HasConfirmed(&_EthDev.CallOpts, _operation, _owner)
}

// HasConfirmed is a free data retrieval call binding the contract method 0xc2cf7326.
//
// Solidity: function hasConfirmed(bytes32 _operation, address _owner) constant returns(bool)
func (_EthDev *EthDevCallerSession) HasConfirmed(_operation [32]byte, _owner common.Address) (bool, error) {
	return _EthDev.Contract.HasConfirmed(&_EthDev.CallOpts, _operation, _owner)
}

// MDailyLimit is a free data retrieval call binding the contract method 0xf1736d86.
//
// Solidity: function m_dailyLimit() constant returns(uint256)
func (_EthDev *EthDevCaller) MDailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthDev.contract.Call(opts, out, "m_dailyLimit")
	return *ret0, err
}

// MDailyLimit is a free data retrieval call binding the contract method 0xf1736d86.
//
// Solidity: function m_dailyLimit() constant returns(uint256)
func (_EthDev *EthDevSession) MDailyLimit() (*big.Int, error) {
	return _EthDev.Contract.MDailyLimit(&_EthDev.CallOpts)
}

// MDailyLimit is a free data retrieval call binding the contract method 0xf1736d86.
//
// Solidity: function m_dailyLimit() constant returns(uint256)
func (_EthDev *EthDevCallerSession) MDailyLimit() (*big.Int, error) {
	return _EthDev.Contract.MDailyLimit(&_EthDev.CallOpts)
}

// MNumOwners is a free data retrieval call binding the contract method 0x4123cb6b.
//
// Solidity: function m_numOwners() constant returns(uint256)
func (_EthDev *EthDevCaller) MNumOwners(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthDev.contract.Call(opts, out, "m_numOwners")
	return *ret0, err
}

// MNumOwners is a free data retrieval call binding the contract method 0x4123cb6b.
//
// Solidity: function m_numOwners() constant returns(uint256)
func (_EthDev *EthDevSession) MNumOwners() (*big.Int, error) {
	return _EthDev.Contract.MNumOwners(&_EthDev.CallOpts)
}

// MNumOwners is a free data retrieval call binding the contract method 0x4123cb6b.
//
// Solidity: function m_numOwners() constant returns(uint256)
func (_EthDev *EthDevCallerSession) MNumOwners() (*big.Int, error) {
	return _EthDev.Contract.MNumOwners(&_EthDev.CallOpts)
}

// MRequired is a free data retrieval call binding the contract method 0x746c9171.
//
// Solidity: function m_required() constant returns(uint256)
func (_EthDev *EthDevCaller) MRequired(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthDev.contract.Call(opts, out, "m_required")
	return *ret0, err
}

// MRequired is a free data retrieval call binding the contract method 0x746c9171.
//
// Solidity: function m_required() constant returns(uint256)
func (_EthDev *EthDevSession) MRequired() (*big.Int, error) {
	return _EthDev.Contract.MRequired(&_EthDev.CallOpts)
}

// MRequired is a free data retrieval call binding the contract method 0x746c9171.
//
// Solidity: function m_required() constant returns(uint256)
func (_EthDev *EthDevCallerSession) MRequired() (*big.Int, error) {
	return _EthDev.Contract.MRequired(&_EthDev.CallOpts)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address _owner) returns()
func (_EthDev *EthDevTransactor) AddOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "addOwner", _owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address _owner) returns()
func (_EthDev *EthDevSession) AddOwner(_owner common.Address) (*types.Transaction, error) {
	return _EthDev.Contract.AddOwner(&_EthDev.TransactOpts, _owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address _owner) returns()
func (_EthDev *EthDevTransactorSession) AddOwner(_owner common.Address) (*types.Transaction, error) {
	return _EthDev.Contract.AddOwner(&_EthDev.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address _from, address _to) returns()
func (_EthDev *EthDevTransactor) ChangeOwner(opts *bind.TransactOpts, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "changeOwner", _from, _to)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address _from, address _to) returns()
func (_EthDev *EthDevSession) ChangeOwner(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _EthDev.Contract.ChangeOwner(&_EthDev.TransactOpts, _from, _to)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address _from, address _to) returns()
func (_EthDev *EthDevTransactorSession) ChangeOwner(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _EthDev.Contract.ChangeOwner(&_EthDev.TransactOpts, _from, _to)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _newRequired) returns()
func (_EthDev *EthDevTransactor) ChangeRequirement(opts *bind.TransactOpts, _newRequired *big.Int) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "changeRequirement", _newRequired)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _newRequired) returns()
func (_EthDev *EthDevSession) ChangeRequirement(_newRequired *big.Int) (*types.Transaction, error) {
	return _EthDev.Contract.ChangeRequirement(&_EthDev.TransactOpts, _newRequired)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _newRequired) returns()
func (_EthDev *EthDevTransactorSession) ChangeRequirement(_newRequired *big.Int) (*types.Transaction, error) {
	return _EthDev.Contract.ChangeRequirement(&_EthDev.TransactOpts, _newRequired)
}

// Confirm is a paid mutator transaction binding the contract method 0x797af627.
//
// Solidity: function confirm(bytes32 _h) returns(bool)
func (_EthDev *EthDevTransactor) Confirm(opts *bind.TransactOpts, _h [32]byte) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "confirm", _h)
}

// Confirm is a paid mutator transaction binding the contract method 0x797af627.
//
// Solidity: function confirm(bytes32 _h) returns(bool)
func (_EthDev *EthDevSession) Confirm(_h [32]byte) (*types.Transaction, error) {
	return _EthDev.Contract.Confirm(&_EthDev.TransactOpts, _h)
}

// Confirm is a paid mutator transaction binding the contract method 0x797af627.
//
// Solidity: function confirm(bytes32 _h) returns(bool)
func (_EthDev *EthDevTransactorSession) Confirm(_h [32]byte) (*types.Transaction, error) {
	return _EthDev.Contract.Confirm(&_EthDev.TransactOpts, _h)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bytes32 _r)
func (_EthDev *EthDevTransactor) Execute(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "execute", _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bytes32 _r)
func (_EthDev *EthDevSession) Execute(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _EthDev.Contract.Execute(&_EthDev.TransactOpts, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bytes32 _r)
func (_EthDev *EthDevTransactorSession) Execute(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _EthDev.Contract.Execute(&_EthDev.TransactOpts, _to, _value, _data)
}

// IsOwner is a paid mutator transaction binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _addr) returns(bool)
func (_EthDev *EthDevTransactor) IsOwner(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "isOwner", _addr)
}

// IsOwner is a paid mutator transaction binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _addr) returns(bool)
func (_EthDev *EthDevSession) IsOwner(_addr common.Address) (*types.Transaction, error) {
	return _EthDev.Contract.IsOwner(&_EthDev.TransactOpts, _addr)
}

// IsOwner is a paid mutator transaction binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _addr) returns(bool)
func (_EthDev *EthDevTransactorSession) IsOwner(_addr common.Address) (*types.Transaction, error) {
	return _EthDev.Contract.IsOwner(&_EthDev.TransactOpts, _addr)
}

// Kill is a paid mutator transaction binding the contract method 0xcbf0b0c0.
//
// Solidity: function kill(address _to) returns()
func (_EthDev *EthDevTransactor) Kill(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "kill", _to)
}

// Kill is a paid mutator transaction binding the contract method 0xcbf0b0c0.
//
// Solidity: function kill(address _to) returns()
func (_EthDev *EthDevSession) Kill(_to common.Address) (*types.Transaction, error) {
	return _EthDev.Contract.Kill(&_EthDev.TransactOpts, _to)
}

// Kill is a paid mutator transaction binding the contract method 0xcbf0b0c0.
//
// Solidity: function kill(address _to) returns()
func (_EthDev *EthDevTransactorSession) Kill(_to common.Address) (*types.Transaction, error) {
	return _EthDev.Contract.Kill(&_EthDev.TransactOpts, _to)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address _owner) returns()
func (_EthDev *EthDevTransactor) RemoveOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "removeOwner", _owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address _owner) returns()
func (_EthDev *EthDevSession) RemoveOwner(_owner common.Address) (*types.Transaction, error) {
	return _EthDev.Contract.RemoveOwner(&_EthDev.TransactOpts, _owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address _owner) returns()
func (_EthDev *EthDevTransactorSession) RemoveOwner(_owner common.Address) (*types.Transaction, error) {
	return _EthDev.Contract.RemoveOwner(&_EthDev.TransactOpts, _owner)
}

// ResetSpentToday is a paid mutator transaction binding the contract method 0x5c52c2f5.
//
// Solidity: function resetSpentToday() returns()
func (_EthDev *EthDevTransactor) ResetSpentToday(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "resetSpentToday")
}

// ResetSpentToday is a paid mutator transaction binding the contract method 0x5c52c2f5.
//
// Solidity: function resetSpentToday() returns()
func (_EthDev *EthDevSession) ResetSpentToday() (*types.Transaction, error) {
	return _EthDev.Contract.ResetSpentToday(&_EthDev.TransactOpts)
}

// ResetSpentToday is a paid mutator transaction binding the contract method 0x5c52c2f5.
//
// Solidity: function resetSpentToday() returns()
func (_EthDev *EthDevTransactorSession) ResetSpentToday() (*types.Transaction, error) {
	return _EthDev.Contract.ResetSpentToday(&_EthDev.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 _operation) returns()
func (_EthDev *EthDevTransactor) Revoke(opts *bind.TransactOpts, _operation [32]byte) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "revoke", _operation)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 _operation) returns()
func (_EthDev *EthDevSession) Revoke(_operation [32]byte) (*types.Transaction, error) {
	return _EthDev.Contract.Revoke(&_EthDev.TransactOpts, _operation)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 _operation) returns()
func (_EthDev *EthDevTransactorSession) Revoke(_operation [32]byte) (*types.Transaction, error) {
	return _EthDev.Contract.Revoke(&_EthDev.TransactOpts, _operation)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 _newLimit) returns()
func (_EthDev *EthDevTransactor) SetDailyLimit(opts *bind.TransactOpts, _newLimit *big.Int) (*types.Transaction, error) {
	return _EthDev.contract.Transact(opts, "setDailyLimit", _newLimit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 _newLimit) returns()
func (_EthDev *EthDevSession) SetDailyLimit(_newLimit *big.Int) (*types.Transaction, error) {
	return _EthDev.Contract.SetDailyLimit(&_EthDev.TransactOpts, _newLimit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 _newLimit) returns()
func (_EthDev *EthDevTransactorSession) SetDailyLimit(_newLimit *big.Int) (*types.Transaction, error) {
	return _EthDev.Contract.SetDailyLimit(&_EthDev.TransactOpts, _newLimit)
}

// EthDevConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the EthDev contract.
type EthDevConfirmationIterator struct {
	Event *EthDevConfirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthDevConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDevConfirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthDevConfirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthDevConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDevConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDevConfirmation represents a Confirmation event raised by the EthDev contract.
type EthDevConfirmation struct {
	Owner     common.Address
	Operation [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0xe1c52dc63b719ade82e8bea94cc41a0d5d28e4aaf536adb5e9cccc9ff8c1aeda.
//
// Solidity: event Confirmation(address owner, bytes32 operation)
func (_EthDev *EthDevFilterer) FilterConfirmation(opts *bind.FilterOpts) (*EthDevConfirmationIterator, error) {

	logs, sub, err := _EthDev.contract.FilterLogs(opts, "Confirmation")
	if err != nil {
		return nil, err
	}
	return &EthDevConfirmationIterator{contract: _EthDev.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0xe1c52dc63b719ade82e8bea94cc41a0d5d28e4aaf536adb5e9cccc9ff8c1aeda.
//
// Solidity: event Confirmation(address owner, bytes32 operation)
func (_EthDev *EthDevFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *EthDevConfirmation) (event.Subscription, error) {

	logs, sub, err := _EthDev.contract.WatchLogs(opts, "Confirmation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDevConfirmation)
				if err := _EthDev.contract.UnpackLog(event, "Confirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EthDevConfirmationNeededIterator is returned from FilterConfirmationNeeded and is used to iterate over the raw logs and unpacked data for ConfirmationNeeded events raised by the EthDev contract.
type EthDevConfirmationNeededIterator struct {
	Event *EthDevConfirmationNeeded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthDevConfirmationNeededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDevConfirmationNeeded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthDevConfirmationNeeded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthDevConfirmationNeededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDevConfirmationNeededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDevConfirmationNeeded represents a ConfirmationNeeded event raised by the EthDev contract.
type EthDevConfirmationNeeded struct {
	Operation [32]byte
	Initiator common.Address
	Value     *big.Int
	To        common.Address
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmationNeeded is a free log retrieval operation binding the contract event 0x1733cbb53659d713b79580f79f3f9ff215f78a7c7aa45890f3b89fc5cddfbf32.
//
// Solidity: event ConfirmationNeeded(bytes32 operation, address initiator, uint256 value, address to, bytes data)
func (_EthDev *EthDevFilterer) FilterConfirmationNeeded(opts *bind.FilterOpts) (*EthDevConfirmationNeededIterator, error) {

	logs, sub, err := _EthDev.contract.FilterLogs(opts, "ConfirmationNeeded")
	if err != nil {
		return nil, err
	}
	return &EthDevConfirmationNeededIterator{contract: _EthDev.contract, event: "ConfirmationNeeded", logs: logs, sub: sub}, nil
}

// WatchConfirmationNeeded is a free log subscription operation binding the contract event 0x1733cbb53659d713b79580f79f3f9ff215f78a7c7aa45890f3b89fc5cddfbf32.
//
// Solidity: event ConfirmationNeeded(bytes32 operation, address initiator, uint256 value, address to, bytes data)
func (_EthDev *EthDevFilterer) WatchConfirmationNeeded(opts *bind.WatchOpts, sink chan<- *EthDevConfirmationNeeded) (event.Subscription, error) {

	logs, sub, err := _EthDev.contract.WatchLogs(opts, "ConfirmationNeeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDevConfirmationNeeded)
				if err := _EthDev.contract.UnpackLog(event, "ConfirmationNeeded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EthDevDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EthDev contract.
type EthDevDepositIterator struct {
	Event *EthDevDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthDevDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDevDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthDevDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthDevDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDevDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDevDeposit represents a Deposit event raised by the EthDev contract.
type EthDevDeposit struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address _from, uint256 value)
func (_EthDev *EthDevFilterer) FilterDeposit(opts *bind.FilterOpts) (*EthDevDepositIterator, error) {

	logs, sub, err := _EthDev.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &EthDevDepositIterator{contract: _EthDev.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address _from, uint256 value)
func (_EthDev *EthDevFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EthDevDeposit) (event.Subscription, error) {

	logs, sub, err := _EthDev.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDevDeposit)
				if err := _EthDev.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EthDevMultiTransactIterator is returned from FilterMultiTransact and is used to iterate over the raw logs and unpacked data for MultiTransact events raised by the EthDev contract.
type EthDevMultiTransactIterator struct {
	Event *EthDevMultiTransact // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthDevMultiTransactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDevMultiTransact)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthDevMultiTransact)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthDevMultiTransactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDevMultiTransactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDevMultiTransact represents a MultiTransact event raised by the EthDev contract.
type EthDevMultiTransact struct {
	Owner     common.Address
	Operation [32]byte
	Value     *big.Int
	To        common.Address
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMultiTransact is a free log retrieval operation binding the contract event 0xe7c957c06e9a662c1a6c77366179f5b702b97651dc28eee7d5bf1dff6e40bb4a.
//
// Solidity: event MultiTransact(address owner, bytes32 operation, uint256 value, address to, bytes data)
func (_EthDev *EthDevFilterer) FilterMultiTransact(opts *bind.FilterOpts) (*EthDevMultiTransactIterator, error) {

	logs, sub, err := _EthDev.contract.FilterLogs(opts, "MultiTransact")
	if err != nil {
		return nil, err
	}
	return &EthDevMultiTransactIterator{contract: _EthDev.contract, event: "MultiTransact", logs: logs, sub: sub}, nil
}

// WatchMultiTransact is a free log subscription operation binding the contract event 0xe7c957c06e9a662c1a6c77366179f5b702b97651dc28eee7d5bf1dff6e40bb4a.
//
// Solidity: event MultiTransact(address owner, bytes32 operation, uint256 value, address to, bytes data)
func (_EthDev *EthDevFilterer) WatchMultiTransact(opts *bind.WatchOpts, sink chan<- *EthDevMultiTransact) (event.Subscription, error) {

	logs, sub, err := _EthDev.contract.WatchLogs(opts, "MultiTransact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDevMultiTransact)
				if err := _EthDev.contract.UnpackLog(event, "MultiTransact", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EthDevOwnerAddedIterator is returned from FilterOwnerAdded and is used to iterate over the raw logs and unpacked data for OwnerAdded events raised by the EthDev contract.
type EthDevOwnerAddedIterator struct {
	Event *EthDevOwnerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthDevOwnerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDevOwnerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthDevOwnerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthDevOwnerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDevOwnerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDevOwnerAdded represents a OwnerAdded event raised by the EthDev contract.
type EthDevOwnerAdded struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerAdded is a free log retrieval operation binding the contract event 0x994a936646fe87ffe4f1e469d3d6aa417d6b855598397f323de5b449f765f0c3.
//
// Solidity: event OwnerAdded(address newOwner)
func (_EthDev *EthDevFilterer) FilterOwnerAdded(opts *bind.FilterOpts) (*EthDevOwnerAddedIterator, error) {

	logs, sub, err := _EthDev.contract.FilterLogs(opts, "OwnerAdded")
	if err != nil {
		return nil, err
	}
	return &EthDevOwnerAddedIterator{contract: _EthDev.contract, event: "OwnerAdded", logs: logs, sub: sub}, nil
}

// WatchOwnerAdded is a free log subscription operation binding the contract event 0x994a936646fe87ffe4f1e469d3d6aa417d6b855598397f323de5b449f765f0c3.
//
// Solidity: event OwnerAdded(address newOwner)
func (_EthDev *EthDevFilterer) WatchOwnerAdded(opts *bind.WatchOpts, sink chan<- *EthDevOwnerAdded) (event.Subscription, error) {

	logs, sub, err := _EthDev.contract.WatchLogs(opts, "OwnerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDevOwnerAdded)
				if err := _EthDev.contract.UnpackLog(event, "OwnerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EthDevOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the EthDev contract.
type EthDevOwnerChangedIterator struct {
	Event *EthDevOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthDevOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDevOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthDevOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthDevOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDevOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDevOwnerChanged represents a OwnerChanged event raised by the EthDev contract.
type EthDevOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_EthDev *EthDevFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*EthDevOwnerChangedIterator, error) {

	logs, sub, err := _EthDev.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &EthDevOwnerChangedIterator{contract: _EthDev.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_EthDev *EthDevFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *EthDevOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _EthDev.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDevOwnerChanged)
				if err := _EthDev.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EthDevOwnerRemovedIterator is returned from FilterOwnerRemoved and is used to iterate over the raw logs and unpacked data for OwnerRemoved events raised by the EthDev contract.
type EthDevOwnerRemovedIterator struct {
	Event *EthDevOwnerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthDevOwnerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDevOwnerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthDevOwnerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthDevOwnerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDevOwnerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDevOwnerRemoved represents a OwnerRemoved event raised by the EthDev contract.
type EthDevOwnerRemoved struct {
	OldOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoved is a free log retrieval operation binding the contract event 0x58619076adf5bb0943d100ef88d52d7c3fd691b19d3a9071b555b651fbf418da.
//
// Solidity: event OwnerRemoved(address oldOwner)
func (_EthDev *EthDevFilterer) FilterOwnerRemoved(opts *bind.FilterOpts) (*EthDevOwnerRemovedIterator, error) {

	logs, sub, err := _EthDev.contract.FilterLogs(opts, "OwnerRemoved")
	if err != nil {
		return nil, err
	}
	return &EthDevOwnerRemovedIterator{contract: _EthDev.contract, event: "OwnerRemoved", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoved is a free log subscription operation binding the contract event 0x58619076adf5bb0943d100ef88d52d7c3fd691b19d3a9071b555b651fbf418da.
//
// Solidity: event OwnerRemoved(address oldOwner)
func (_EthDev *EthDevFilterer) WatchOwnerRemoved(opts *bind.WatchOpts, sink chan<- *EthDevOwnerRemoved) (event.Subscription, error) {

	logs, sub, err := _EthDev.contract.WatchLogs(opts, "OwnerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDevOwnerRemoved)
				if err := _EthDev.contract.UnpackLog(event, "OwnerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EthDevRequirementChangedIterator is returned from FilterRequirementChanged and is used to iterate over the raw logs and unpacked data for RequirementChanged events raised by the EthDev contract.
type EthDevRequirementChangedIterator struct {
	Event *EthDevRequirementChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthDevRequirementChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDevRequirementChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthDevRequirementChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthDevRequirementChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDevRequirementChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDevRequirementChanged represents a RequirementChanged event raised by the EthDev contract.
type EthDevRequirementChanged struct {
	NewRequirement *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRequirementChanged is a free log retrieval operation binding the contract event 0xacbdb084c721332ac59f9b8e392196c9eb0e4932862da8eb9beaf0dad4f550da.
//
// Solidity: event RequirementChanged(uint256 newRequirement)
func (_EthDev *EthDevFilterer) FilterRequirementChanged(opts *bind.FilterOpts) (*EthDevRequirementChangedIterator, error) {

	logs, sub, err := _EthDev.contract.FilterLogs(opts, "RequirementChanged")
	if err != nil {
		return nil, err
	}
	return &EthDevRequirementChangedIterator{contract: _EthDev.contract, event: "RequirementChanged", logs: logs, sub: sub}, nil
}

// WatchRequirementChanged is a free log subscription operation binding the contract event 0xacbdb084c721332ac59f9b8e392196c9eb0e4932862da8eb9beaf0dad4f550da.
//
// Solidity: event RequirementChanged(uint256 newRequirement)
func (_EthDev *EthDevFilterer) WatchRequirementChanged(opts *bind.WatchOpts, sink chan<- *EthDevRequirementChanged) (event.Subscription, error) {

	logs, sub, err := _EthDev.contract.WatchLogs(opts, "RequirementChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDevRequirementChanged)
				if err := _EthDev.contract.UnpackLog(event, "RequirementChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EthDevRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the EthDev contract.
type EthDevRevokeIterator struct {
	Event *EthDevRevoke // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthDevRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDevRevoke)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthDevRevoke)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthDevRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDevRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDevRevoke represents a Revoke event raised by the EthDev contract.
type EthDevRevoke struct {
	Owner     common.Address
	Operation [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0xc7fb647e59b18047309aa15aad418e5d7ca96d173ad704f1031a2c3d7591734b.
//
// Solidity: event Revoke(address owner, bytes32 operation)
func (_EthDev *EthDevFilterer) FilterRevoke(opts *bind.FilterOpts) (*EthDevRevokeIterator, error) {

	logs, sub, err := _EthDev.contract.FilterLogs(opts, "Revoke")
	if err != nil {
		return nil, err
	}
	return &EthDevRevokeIterator{contract: _EthDev.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0xc7fb647e59b18047309aa15aad418e5d7ca96d173ad704f1031a2c3d7591734b.
//
// Solidity: event Revoke(address owner, bytes32 operation)
func (_EthDev *EthDevFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *EthDevRevoke) (event.Subscription, error) {

	logs, sub, err := _EthDev.contract.WatchLogs(opts, "Revoke")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDevRevoke)
				if err := _EthDev.contract.UnpackLog(event, "Revoke", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EthDevSingleTransactIterator is returned from FilterSingleTransact and is used to iterate over the raw logs and unpacked data for SingleTransact events raised by the EthDev contract.
type EthDevSingleTransactIterator struct {
	Event *EthDevSingleTransact // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthDevSingleTransactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDevSingleTransact)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthDevSingleTransact)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthDevSingleTransactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDevSingleTransactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDevSingleTransact represents a SingleTransact event raised by the EthDev contract.
type EthDevSingleTransact struct {
	Owner common.Address
	Value *big.Int
	To    common.Address
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSingleTransact is a free log retrieval operation binding the contract event 0x92ca3a80853e6663fa31fa10b99225f18d4902939b4c53a9caae9043f6efd004.
//
// Solidity: event SingleTransact(address owner, uint256 value, address to, bytes data)
func (_EthDev *EthDevFilterer) FilterSingleTransact(opts *bind.FilterOpts) (*EthDevSingleTransactIterator, error) {

	logs, sub, err := _EthDev.contract.FilterLogs(opts, "SingleTransact")
	if err != nil {
		return nil, err
	}
	return &EthDevSingleTransactIterator{contract: _EthDev.contract, event: "SingleTransact", logs: logs, sub: sub}, nil
}

// WatchSingleTransact is a free log subscription operation binding the contract event 0x92ca3a80853e6663fa31fa10b99225f18d4902939b4c53a9caae9043f6efd004.
//
// Solidity: event SingleTransact(address owner, uint256 value, address to, bytes data)
func (_EthDev *EthDevFilterer) WatchSingleTransact(opts *bind.WatchOpts, sink chan<- *EthDevSingleTransact) (event.Subscription, error) {

	logs, sub, err := _EthDev.contract.WatchLogs(opts, "SingleTransact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDevSingleTransact)
				if err := _EthDev.contract.UnpackLog(event, "SingleTransact", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
