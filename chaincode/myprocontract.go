package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

//SmartContract empty structure
type SmartContract struct {
}

// User Structure.  Structure tags are used by encoding/json library
type User struct {
	Objectype string `json:"doctype"`
	ID        string `json:"Id"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Address1  string `json:"Address1"`
	Address2  string `json:"Address2"`
	City      string `json:"City"`
	State     string `json:"State"`
	Zipcode   string `json:"Zipcode"`
	Country   string `json:"Country"`
	Email     string `json:"Email"`
	Phone     string `json:"Phone"`
	Usertype  string `json:"Usertype"`
	Loginid   string `json:"Loginid"`
	Password  string `json:"Password"`
}

// Bank structure. Structure tags are used by encoding/json library
type Bank struct {
	Objectype   string `json:"doctype"`
	ID          string `json:"Id"`
	Loginid     string `json:"Loginid"`
	Profilename string `json:"Profilename"`
	Accountype  string `json:"Accountype"`
	Bankname    string `json:"Bankname"`
	Routing     string `json:"Routing"`
	Account     string `json:"Account"`
	Validated   string `json:"Validated"`
	Sharedlogin string `json:"Sharedlogin"`
}

/*Init Function bypassed
Init.. method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
Best practice is to have any Ledger initialization in separate function -- see initLedger()
Function bypassed by Mahender as per best practice
*/
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	//func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("MyPro Is Starting Up")
	funcName, args := APIstub.GetFunctionAndParameters()
	var number int
	var err error
	txID := APIstub.GetTxID()

	fmt.Println("Init() is running")
	fmt.Println("Transaction ID:", txID)
	fmt.Println("  GetFunctionAndParameters() function:", funcName)
	fmt.Println("  GetFunctionAndParameters() args count:", len(args))
	fmt.Println("  GetFunctionAndParameters() args found:", args)

	// expecting 1 arg for instantiate or upgrade
	if len(args) == 1 {
		fmt.Println("  GetFunctionAndParameters() arg[0] length", len(args[0]))

		// expecting arg[0] to be length 0 for upgrade
		if len(args[0]) == 0 {
			fmt.Println("  Uh oh, args[0] is empty...")
		} else {
			fmt.Println("  Great news everyone, args[0] is not empty")

			// convert numeric string to integer
			number, err = strconv.Atoi(args[0])
			if err != nil {
				return shim.Error("Expecting a numeric string argument to Init() for instantiate")
			}

			// this is a very simple test. let's write to the ledger and error out on any errors
			// it's handy to read this right away to verify network is healthy if it wrote the correct value
			err = APIstub.PutState("selftest", []byte(strconv.Itoa(number)))
			if err != nil {
				return shim.Error(err.Error()) //self-test fail
			}
		}
	}

	// showing the alternative argument shim function
	alt := APIstub.GetStringArgs()
	fmt.Println("  GetStringArgs() args count:", len(alt))
	fmt.Println("  GetStringArgs() args found:", alt)

	// store compatible mypro application version
	err = APIstub.PutState("mypro_ui", []byte("4.0.1"))
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("Ready for action") //self-test pass
	return shim.Success(nil)
}

/*Invoke function bypassed
 * Invoke.. method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryUser" {
		fmt.Println("Calling queryUser from main invoke")
		return s.queryUser(APIstub, args)
	} else if function == "createUser" {
		fmt.Println("calling createUser from main invoke")
		return createUser(APIstub, args)
	} else if function == "createBank" {
		fmt.Println("Calling createBank from main invoke")
		return s.createBank(APIstub, args)
	} else if function == "queryBank" {
		fmt.Println("Calling query bank from main invoke")
		return s.queryBank(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	fmt.Println("inside smart contract with user id for user", args[0])
	fmt.Println("inside smart contract var count for user", len(args))
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var user User
	user.Loginid = args[0]
	key := "U" + args[0]
	fmt.Println("value of user.LoginID and key is", user.Loginid, key)
	//userAsBytes, _ := APIstub.GetState(args[0])
	userAsBytes, _ := APIstub.GetState(key)
	fmt.Println("inside smart contract user bytes are", userAsBytes)
	return shim.Success(userAsBytes)
}

func (s *SmartContract) queryBank(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	fmt.Println("inside smart contract with user id for Bank ", args[0])
	fmt.Println("inside smart contract var count for bank", len(args))
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var bank Bank
	bank.Loginid = args[0]
	key := "B" + args[0]
	fmt.Println("value of bank.LoginID and key is", bank.Loginid, key)
	//bankAsBytes, _ := APIstub.GetState(args[0])
	bankAsBytes, _ := APIstub.GetState(key)
	fmt.Println("inside smart contract bank bytes are", bankAsBytes)
	return shim.Success(bankAsBytes)
}

func createUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	fmt.Println("starting createUser")
	fmt.Println("Number of arguments", len(args))
	if len(args) != 14 {
		return shim.Error("Incorrect number of arguments. Expecting 14")
	}
	var user User
	user.Objectype = "mypro_users"
	user.ID = args[0]
	user.Firstname = args[1]
	user.Lastname = args[2]
	user.Address1 = args[3]
	user.Address2 = args[4]
	user.City = args[5]
	user.State = args[6]
	user.Zipcode = args[7]
	user.Country = args[8]
	user.Email = args[9]
	user.Phone = args[10]
	user.Usertype = args[11]
	user.Loginid = args[12]
	user.Password = args[13]
	fmt.Println("User is", user)
	/*check if user already exists
	_, err = getOwner(APIstub, user.ID)
	if err == nil {
		fmt.Println("This owner already exists - " + user.ID)
		return shim.Error("This owner already exists - " + user.ID)
	}
	*/
	userAsBytes, _ := json.Marshal(user)
	fmt.Println("user bytes are", userAsBytes)
	key := "U" + args[12]
	fmt.Println("User key is", key)
	err = APIstub.PutState(key, userAsBytes)
	//err = APIstub.PutState(args[12], userAsBytes)
	if err != nil {
		fmt.Println("Could not store user")
		return shim.Error(err.Error())
	}
	fmt.Println("- end createUser mypro")
	return shim.Success(nil)

}

func (s *SmartContract) createBank(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	fmt.Println("starting createBank")
	fmt.Println("Number of bank arguments", len(args))
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments for bank. Expecting 9")
	}
	var bank Bank
	bank.Objectype = "mypro_bank"
	bank.ID = args[0]
	bank.Loginid = args[1]
	bank.Profilename = args[2]
	bank.Accountype = args[3]
	bank.Bankname = args[4]
	bank.Routing = args[5]
	bank.Account = args[6]
	bank.Validated = args[7]
	bank.Sharedlogin = args[8]
	fmt.Println("Bank is", bank)
	bankAsBytes, _ := json.Marshal(bank)
	fmt.Println("bank bytes are", bankAsBytes)
	key := "B" + args[1]
	fmt.Println("key for bank is", key)
	err = APIstub.PutState(key, bankAsBytes)
	//err = APIstub.PutState(args[1], bankAsBytes)
	if err != nil {
		fmt.Println("Could not store bank")
		return shim.Error(err.Error())
	}
	fmt.Println("- end createBank mypro")
	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
