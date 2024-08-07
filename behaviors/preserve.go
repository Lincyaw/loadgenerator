package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	// Preserve - Main
	AccountID  = "accountId"
	ContactsID = "contactsId"
	TripID     = "tripId"
	SeatType   = "seatType"
	//LoginToken = "loginToken"
	Date            = "date"
	From            = "from"
	To              = "to"
	Assurance       = "assurance"
	FoodType        = "foodType"
	StationName     = "stationName"
	StoreName       = "storeName"
	FoodName        = "foodName"
	FoodPrice       = "foodPrice"
	HandleDate      = "handleDate"
	ConsigneeName   = "consigneeName"
	ConsigneePhone  = "consigneePhone"
	ConsigneeWeight = "consigneeWeight"
	IsWithin        = "isWithin"

	// Assurance
	OrderId   = "orderId"
	TypeIndex = "typeIndex"
	TypeName  = "typeName"
	TypePrice = "typePrice"

	// VerifyCode
	BooleanVerifyCode = "booleanVerifyCode"

	// User
	UserID       = "userId"
	UserName     = "userName"
	Password     = "password"
	Gender       = "gender"
	DocumentType = "documentType"
	DocumentNum  = "documentNum"
	Email        = "email"

	// Contacts
	//Id = "id" - ContactsID
	//AccountId = "accountId"
	Name = "name"
	//DocumentType   = "documentType"
	DocumentNumber = "documentNumber"
	PhoneNumber    = "phoneNumber"

	// Consign
	ID      = "id"
	OrderID = "orderId"
	//AccountID = "accountId"
	//HandleDate = "handleDate"
	TargetDate = "targetDate"
	//From = "from"
	//To = "to"
	Consignee = "consignee"
	Phone     = "phone"
	Weight    = "weight"
	//IsWithin = "isWithin"
	Price = "price"

	// FoodBehavior

	// Route
	RouteID = "routeId"
)

var PreserveBehaviorChain *Chain

func init() {
	// ------------------------------------- init -------------------------------------------
	// ------------------------------------- init -------------------------------------------
	// ------------------------------------- init -------------------------------------------
	// Main Chain
	PreserveBehaviorChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("PreserveBehaviors(Chain) Statrs. Starts time: %v", time.Now().String())
		return nil, nil
	}, "Dummy"))
	// AssuranceBehaviorChain -
	AssuranceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("CreateAssuranceChain. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyAssurance"))
	// UserBehaviorsChain
	UserBehaviorsChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("UserBehaviorsChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyUserBehaviors"))
	// VerifyCodeBehaviorChain
	VerifyCodeBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("VerifyCodeBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyVerifyCodeBehavior"))
	// AuthBehaviorChain
	AuthBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("AuthBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyAuthBehavior"))
	// UserBehaviorChain
	UserBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("UserBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyUserBehavior"))
	// ContactsBehaviorChain
	ContactsBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("ContactsBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyContactsBehavior"))
	// ConsignBehaviorsChain
	ConsignBehaviorsChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("ConsignBehaviorsChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyConsignBehaviors"))
	//ConsignPriceBehaviorChain
	ConsignPriceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("ConsignPriceBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyConsignPriceBehavior"))
	// FoodBehaviorChain
	FoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("FoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyFoodBehavior"))
	// TravelBehaviorChain
	TravelBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("TravelBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTravelBehavior"))
	// StationFoodBehaviorChain
	StationFoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("StationFoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationFoodBehavior"))
	//TrainFoodBehaviorChain
	TrainFoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("TrainFoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTrainFoodBehavior"))
	// TrainBehaviorChain
	TrainBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("TrainBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTrainBehavior"))
	//RouteBehaviorChain
	RouteBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("RouteBehaviorChain starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyRouteBehavior"))
	// BasicBehaviorChain
	BasicBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("BasicBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyBasicBehavior"))
	// SeatBehaviorChain
	SeatBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("SeatBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummySeatBehavior"))
	// StationBehaviorChain
	StationBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	//PriceBehaviorChain
	PriceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("PriceBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyPriceBehavior"))
	//ConfigBehaviorChain
	ConfigBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("ConfigBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyConfigBehavior"))
	//OrderBehaviorChain
	OrderBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("OrderBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderBehavior"))
	//OrderOtherBehaviorChain
	OrderOtherBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("OrderOtherBehaviorChain Starts. Strat time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderOtherBehavior"))
	// StationBehaviorChain
	StationBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	// SecurityBehaviorChain
	SecurityBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("SecurityBehaviorChain Satrts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummySecurityBehavior"))

	// ------------------------------------- NewFuncNode -------------------------------------------
	// ------------------------------------- NewFuncNode -------------------------------------------
	// ------------------------------------- NewFuncNode -------------------------------------------
	//AssuranceBehaviorChain - Assurance
	QueryAssuranceNode := NewFuncNode(QueryAssurance, "QueryAssurance")
	// CreateAssuranceNode := NewFuncNode(CreateAssurance, "CreateAssurance")

	//UserBehaviorsChain
	// AuthBehaviorChain - LoginAdmin/LoginBasic
	//LoginAdminNode := NewFuncNode(LoginAdmin, "LoginAdmin")
	LoginBasicNode := NewFuncNode(LoginBasic, "LoginBasic")
	//	VerifyCodeBehaviorChain
	VerifyCodeNode := NewFuncNode(VerifyCode, "VerifyCode")
	// UserBehaviorChain
	QueryUserNode := NewFuncNode(QueryUser, "QueryUser")

	//ContactsBehaviorChain - Contacts
	QueryContactsNode := NewFuncNode(QueryContacts, "QueryContacts")
	CreateContactsNode := NewFuncNode(CreateContacts, "CreateContacts")

	//ConsignBehaviorsChain
	QueryConsignNode := NewFuncNode(QueryConsign, "QueryConsign")
	CreateConsignNode := NewFuncNode(CreateConsign, "CreateConsign")
	// ConsignPriceBehaviorChain
	QueryConsignPriceNode := NewFuncNode(QueryConsignPric, "QueryConsignPrice")
	CreateConsignPriceNode := NewFuncNode(CreateConsignPrice, "CreateConsignPrice")

	//FoodBehaviorChain
	QueryFoodNode := NewFuncNode(QueryFood, "QueryFood")
	// StationFoodBehaviorChain
	QueryStationFoodNode := NewFuncNode(QueryStationFood, "QueryStationFood")
	// TrainFoodBehaviorChain
	QueryTrainFoodNode := NewFuncNode(QueryTrainFood, "QueryTrainFood")
	// TravelBehaviorChain
	QueryTravelNode := NewFuncNode(QueryTrip, "QueryTrip")
	CreateTravelNode := NewFuncNode(CreateTrip, "CreateTrip")

	//TravelBehaviorChain
	// TrainBehaviorChain
	QueryTrainNode := NewFuncNode(QueryTrain, "QueryTrain")
	// RouteBehaviorChain
	QueryRouteNode := NewFuncNode(QueryRoute, "QueryRoute")
	// BasicBehaviorChain
	QueryBasicNode := NewFuncNode(QueryBasic, "QueryBasic")
	// SeatBehaviorChain
	QuerySeatNode := NewFuncNode(QuerySeat, "QuerySeat")

	//BasicBehaviorChain
	// StationBehaviorChain
	QueryStationNode := NewFuncNode(QueryStation, "QueryStation")
	// PriceBehaviorChain
	QueryPriceNode := NewFuncNode(QueryPrice, "QueryPrice")
	// RouteBehaviorChain
	// TrainBehaviorChain

	//SeatBehaviorChain
	// ConfigBehaviorChain
	QueryConfigNode := NewFuncNode(QueryConfig, "QueryConfig")
	// OrderBehaviorChain
	QueryOrderNode := NewFuncNode(QueryOrder, "QueryOrder")
	// OrderOtherBehaviorChain
	QueryOrderOtherNode := NewFuncNode(QueryOrderOther, "QueryOrderOther")

	//OrderBehaviorChain
	// StationBehaviorChain

	//OrderOtherBehaviorChain
	// StationBehaviorChain

	//SecurityBehaviorChain
	// OrderBehaviorChain
	// OrderOtherBehaviorChain

	// ******* Preserve ********
	PreserveNode := NewFuncNode(Preserve, "Preserve")

	// ------------------------------------- NewChain -------------------------------------------
	// ------------------------------------- NewChain -------------------------------------------
	// ------------------------------------- NewChain -------------------------------------------
	// AssuranceBehaviorChain - Assurance
	QueryAssuranceChain := NewChain(QueryAssuranceNode)
	//CreateAssuranceChain := NewChain(CreateAssuranceNode)

	// UserBehaviorsChain
	// AuthBehaviorChain - LoginAdmin/LoginBasic
	LoginBasicChain := NewChain(LoginBasicNode)
	// VerifyCodeBehaviorChain
	VerifyCodeChain := NewChain(VerifyCodeNode)
	// UserBehaviorChain
	QueryUserChain := NewChain(QueryUserNode)

	// ContactsBehaviorChain - Contacts
	QueryContactsChain := NewChain(QueryContactsNode)
	CreateContactsChain := NewChain(CreateContactsNode)

	// ConsignBehaviorsChain
	QueryConsignChain := NewChain(QueryConsignNode)
	CreateConsignChain := NewChain(CreateConsignNode)
	// ConsignPriceBehaviorChain
	QueryConsignPriceChain := NewChain(QueryConsignPriceNode)
	CreateConsignPriceChain := NewChain(CreateConsignPriceNode)

	// FoodBehaviorChain
	QueryFoodChain := NewChain(QueryFoodNode)
	// StationFoodBehaviorChain
	QueryStationFoodChain := NewChain(QueryStationFoodNode)
	// TrainFoodBehaviorChain
	QueryTrainFoodChain := NewChain(QueryTrainFoodNode)
	// TravelBehaviorChain
	QueryTravelChain := NewChain(QueryTravelNode)
	CreateTravelChain := NewChain(CreateTravelNode)

	// TrainBehaviorChain
	QueryTrainChain := NewChain(QueryTrainNode)
	// RouteBehaviorChain
	QueryRouteChain := NewChain(QueryRouteNode)
	// BasicBehaviorChain
	QueryBasicChain := NewChain(QueryBasicNode)
	// SeatBehaviorChain
	QuerySeatChain := NewChain(QuerySeatNode)

	// StationBehaviorChain
	QueryStationChain := NewChain(QueryStationNode)
	// PriceBehaviorChain
	QueryPriceChain := NewChain(QueryPriceNode)

	// ConfigBehaviorChain
	QueryConfigChain := NewChain(QueryConfigNode)
	// OrderBehaviorChain
	QueryOrderChain := NewChain(QueryOrderNode)
	// OrderOtherBehaviorChain
	QueryOrderOtherChain := NewChain(QueryOrderOtherNode)

	// The Last Chain - Preserve Behavior Chain
	PreserveChain := NewChain(PreserveNode)

	// ------------------------------------- AddNextChain -------------------------------------------
	// ------------------------------------- AddNextChain -------------------------------------------
	// ------------------------------------- AddNextChain -------------------------------------------
	// 逆序 - 从处理逆序的第一层开始

	// XXX
	//...

	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	//PreserveChain
	//PreserveChain.AddNextChain(, 1)

	// ------------------------------------- VisualizeChain -------------------------------------------
	fmt.Println(PreserveChain.VisualizeChain(0))
	fmt.Println()
}

// ************************************* NewFuncNode_Function *******************************************
// ************************************* NewFuncNode_Function *******************************************
// ************************************* NewFuncNode_Function *******************************************

// AssuranceBehaviorChain
func QueryAssurance(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	Assurances, err := cli.GetAllAssurances()
	if err != nil {
		log.Fatalf("GetAllAssurances failed: %v", err)
		return nil, err
	}
	if Assurances.Status != 1 {
		log.Fatalf("Assurances status is not 1: %d", Assurances.Status)
		return nil, nil
	}

	randomIndex := rand.Intn(len(Assurances.Data))
	ctx.Set(OrderId, Assurances.Data[randomIndex].OrderId)
	ctx.Set(TypeIndex, Assurances.Data[randomIndex].TypeIndex)
	ctx.Set(TypeName, Assurances.Data[randomIndex].TypeName)
	ctx.Set(TypePrice, Assurances.Data[randomIndex].TypePrice)

	return nil, nil
}

func CreateAssurance(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	//Create a new assurance
	TheOrderID := ctx.Get(OrderId).(string)
	addAssuranceResp, err := cli.CreateNewAssurance(1, TheOrderID) // typeIndex 1 -> TRAFFIC_ACCIDENT
	if err != nil {
		log.Fatalf("CreateNewAssurance failed: %v", err)
		return nil, err
	}
	if addAssuranceResp.Msg == "Already exists" {
		log.Fatalf("Order ID found, skip")
		return nil, err
	}
	if addAssuranceResp.Data.OrderId != TheOrderID {
		log.Fatalf("Request failed, addAssuranceResp.Data.OrderId:%s, expected: %s", addAssuranceResp.Data.OrderId, TheOrderID)
		return nil, err
	}
	if addAssuranceResp.Data.Type != "TRAFFIC_ACCIDENT" {
		log.Fatalf("Request failed, addAssuranceResp.Data.Type are expected to be 'TRAFFIC_ACCIDENT' but actually: %v", addAssuranceResp.Data.Type)
		return nil, err
	}

	ctx.Set(OrderId, addAssuranceResp.Data.OrderId)
	//ctx.Set(TypeIndex, addAssuranceResp.Data.)
	//ctx.Set(TypeName, Assurances.Data[randomIndex].TypeName)
	//ctx.Set(TypePrice, Assurances.Data[randomIndex].TypePrice)

	return nil, nil
}

//UserBehaviorsChain
// LoginBasicChain
//func LoginBasic(ctx *Context) (*NodeResult, error) {
//	cli, ok := ctx.Get(Client).(*service.SvcImpl)
//	if !ok {
//		return nil, fmt.Errorf("service client not found in context")
//	}
//	// login
//	loginResult, err := cli.ReqUserLogin(&service.UserLoginInfoReq{
//		Password:         "111111",
//		UserName:         "fdse_microservice",
//		VerificationCode: "123",
//	})
//	if err != nil {
//		return nil, err
//	}
//	ctx.Set(LoginToken, loginResult.Data.Token)
//	return nil, nil
//}

// VerifyCodeBehaviorChain
func VerifyCode(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	verifyCode := generateVerifyCode()
	verifyCodeResp, err := cli.VerifyCode(verifyCode)
	if err != nil {
		log.Fatalf("Request failed, err %s", err)
		return nil, err
	}
	if !verifyCodeResp {
		log.Fatalf("Verification failed")
		return nil, err
	}
	log.Fatalf("Verification code verified. The result is %v and verifyCode: %v", verifyCodeResp, verifyCode)

	ctx.Set(BooleanVerifyCode, verifyCodeResp)

	return nil, nil
}

func QueryUser(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	allUsersResp, err := cli.GetAllUsers()
	if err != nil {
		log.Fatalf("Request failed, err1 %s", err)
		return nil, err
	}
	if allUsersResp.Status != 1 {
		log.Fatalf("Expected status 200, got %d", allUsersResp.Status)
		return nil, err
	}

	randomIndex := rand.Intn(len(allUsersResp.Data))
	ctx.Set(UserID, allUsersResp.Data[randomIndex].UserID)
	ctx.Set(UserName, allUsersResp.Data[randomIndex].UserName)
	ctx.Set(Password, allUsersResp.Data[randomIndex].Password)
	ctx.Set(Gender, allUsersResp.Data[randomIndex].Gender)
	ctx.Set(DocumentNum, allUsersResp.Data[randomIndex].DocumentNum)
	ctx.Set(DocumentType, allUsersResp.Data[randomIndex].DocumentType)
	ctx.Set(Email, allUsersResp.Data[randomIndex].Email)

	return nil, nil
}

// ContactsBehaviorChain
func QueryContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	GetAllContacts, err := cli.GetAllContacts()
	if err != nil {
		log.Fatalf("[Mock AccountID]GetAllContacts fail. The error occurs: %v", err)
		return nil, err
	}
	if GetAllContacts.Status != 1 {
		log.Fatalf("[Mock AccountID]GetAllContacts.Status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(GetAllContacts.Data))
	ctx.Set(AccountID, GetAllContacts.Data[randomIndex].AccountId)
	ctx.Set(ContactsID, GetAllContacts.Data[randomIndex].Id)
	ctx.Set(Name, GetAllContacts.Data[randomIndex].Name)
	ctx.Set(DocumentType, GetAllContacts.Data[randomIndex].DocumentType)
	ctx.Set(DocumentNumber, GetAllContacts.Data[randomIndex].DocumentNumber)
	ctx.Set(PhoneNumber, GetAllContacts.Data[randomIndex].PhoneNumber)

	return nil, nil
}

func CreateContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	CreateContactsInput := service.AdminContacts{
		Id:             faker.UUIDHyphenated(),
		AccountId:      faker.UUIDHyphenated(),
		Name:           faker.Name(),
		DocumentType:   rand.Intn(1),
		DocumentNumber: generateDocumentNumber(),
		PhoneNumber:    faker.PhoneNumber,
	}
	CreateContacts, err := cli.AddContact(&CreateContactsInput)
	if err != nil {
		log.Fatalf("[Mock AccountID] CreateContacts error occurs: %v", err)
		return nil, err
	}
	if CreateContacts.Status != 1 {
		log.Fatalf("[Mock AccountID] CreateContacts.Status != 1")
		return nil, err
	}

	ctx.Set(AccountID, CreateContacts.Data.AccountId)
	ctx.Set(ContactsID, CreateContacts.Data.Id)
	ctx.Set(Name, CreateContacts.Data.Name)
	ctx.Set(DocumentType, CreateContacts.Data.DocumentType)
	ctx.Set(DocumentNumber, CreateContacts.Data.DocumentNumber)
	ctx.Set(PhoneNumber, CreateContacts.Data.PhoneNumber)

	return nil, nil
}

// ConsignBehaviorsChain
func QueryConsign(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	//TheAccountId := ctx.Get(AccountID).(string)
	//// Query consign records by account ID
	//consignsByAccountId, err := cli.QueryByAccountId(TheAccountId)
	//if err != nil {
	//	log.Fatalf("QueryByAccountId failed: %v", err)
	//}
	//if consignsByAccountId.Status != 1 {
	//	log.Fatalf("consignsByAccountId failed")
	//
	//}
	///*found := false
	//for _, consign := range consignsByAccountId.Data {
	//	if consign.IsWithin == existedConsign.IsWithin &&
	//		consign.To == existedConsign.To &&
	//		consign.Weight == existedConsign.Weight &&
	//		consign.ID == existedConsign.ID &&
	//		consign.Phone == existedConsign.Phone &&
	//		consign.HandleDate == existedConsign.HandleDate &&
	//		consign.TargetDate == existedConsign.TargetDate &&
	//		consign.OrderID == existedConsign.OrderID &&
	//		consign.Consignee == existedConsign.Consignee &&
	//		consign.From == existedConsign.From &&
	//		consign.AccountID == existedConsign.AccountID {
	//		found = true
	//	}
	//}
	//if !found {
	//	log.Fatalf("Can not find consign by accountId.")
	//}*/
	//log.Fatalf("QueryByAccountId response: %+v", consignsByAccountId)

	// Query consign records by order ID
	TheOrderId := ctx.Get(OrderId).(string)
	consignsByOrderId, err := cli.QueryByOrderId(TheOrderId)
	if err != nil {
		log.Fatalf("QueryByOrderId failed: %v", err)
		return nil, err
	}
	if consignsByOrderId.Status != 1 {
		log.Fatalf("consignsByOrderId.Status = 1")
		return nil, err
	}
	/*isMatch1 := false
	if consignsByOrderId.Data.OrderId == existedConsign.OrderID &&
		consignsByOrderId.Data.Id == existedConsign.ID &&
		consignsByOrderId.Data.From == existedConsign.From &&
		consignsByOrderId.Data.To == existedConsign.To &&
		consignsByOrderId.Data.Phone == existedConsign.Phone &&
		consignsByOrderId.Data.Consignee == existedConsign.Consignee &&
		consignsByOrderId.Data.TargetDate == existedConsign.TargetDate &&
		consignsByOrderId.Data.HandleDate == existedConsign.HandleDate &&
		consignsByOrderId.Data.Weight == existedConsign.Weight &&
		consignsByOrderId.Data.AccountId == existedConsign.AccountID {
		isMatch1 = true
	}
	if !isMatch1 {
		log.Fatalf("Can not find consign by orderId.")
	}*/
	//log.Fatalf("QueryByOrderId response: %+v", consignsByOrderId)

	// Query consign records by consignee
	//TheConsignee := ctx.Get(Name).(string)
	//consignsByConsignee, err := cli.QueryByConsignee(TheConsignee)
	//if err != nil {
	//	log.Fatalf("QueryByConsignee failed: %v", err)
	//}
	//if consignsByConsignee.Status != 1 {
	//	log.Fatalf("consignsByConsignee failed.")
	//}
	//isMatch2 := false
	//for _, consign := range consignsByConsignee.Data {
	//	if consign.Id == existedConsign.ID &&
	//		consign.AccountId == existedConsign.AccountID &&
	//		consign.OrderId == existedConsign.OrderID &&
	//		consign.To == existedConsign.To &&
	//		consign.From == existedConsign.From &&
	//		consign.Weight == existedConsign.Weight &&
	//		consign.HandleDate == existedConsign.HandleDate &&
	//		consign.TargetDate == existedConsign.TargetDate &&
	//		consign.Phone == existedConsign.Phone &&
	//		consign.Consignee == existedConsign.Consignee {
	//		isMatch2 = true
	//	}
	//}
	//if !isMatch2 {
	//	log.Fatalf("Can not find consign by consignee.")
	//}
	//log.Fatalf("QueryByConsignee response: %+v", consignsByConsignee)

	ctx.Set(ID, consignsByOrderId.Data.Id)
	ctx.Set(OrderId, consignsByOrderId.Data.OrderId)
	ctx.Set(AccountID, consignsByOrderId.Data.AccountId)
	ctx.Set(HandleDate, consignsByOrderId.Data.HandleDate)
	ctx.Set(TargetDate, consignsByOrderId.Data.TargetDate)
	ctx.Set(From, consignsByOrderId.Data.From)
	ctx.Set(To, consignsByOrderId.Data.To)
	ctx.Set(Consignee, consignsByOrderId.Data.Consignee)
	ctx.Set(Phone, consignsByOrderId.Data.Phone)
	ctx.Set(Weight, consignsByOrderId.Data.Weight)
	ctx.Set(Price, consignsByOrderId.Data.Price)

	return nil, nil
}

func CreateConsign(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock data
	MockedId := faker.UUIDHyphenated()
	MockedAccountId := ctx.Get(AccountID).(string)
	MockedOrderId := ctx.Get(OrderId).(string)
	MockedHandleDate := ctx.Get(HandleDate).(string)
	MockedTargetDate := ctx.Get(TargetDate).(string)
	MockedFromPlace := ctx.Get(From).(string)
	MockedToPlace := ctx.Get(To).(string)
	MockedConsignee := ctx.Get(ConsigneeName).(string)
	MockedPhone := ctx.Get(PhoneNumber).(string)
	MockedWeight := GenerateWeight()

	// Insert a new consign record
	insertReq := service.Consign{
		ID:         MockedId,
		OrderID:    MockedOrderId,
		AccountID:  MockedAccountId,
		HandleDate: MockedHandleDate,
		TargetDate: MockedTargetDate,
		From:       MockedFromPlace,
		To:         MockedToPlace,
		Consignee:  MockedConsignee,
		Phone:      MockedPhone,
		Weight:     MockedWeight,
		IsWithin:   BooleanIsWithin(MockedWeight),
	}
	insertResp, err := cli.InsertConsignRecord(&insertReq)
	if err != nil {
		log.Fatalf("InsertConsignRecord failed: %v", err)
		return nil, err
	}
	if insertResp.Msg == "Already exists" {
		return nil, fmt.Errorf("Consign already exists")
	}
	if insertResp.Status != 1 {
		log.Fatalf("InsertConsignRecord failed: %v", insertResp.Status)
		return nil, err
	}
	isMatch := false
	if /*insertResp.Data.ID == insertReq.ID &&*/
	insertResp.Data.IsWithin == insertReq.IsWithin &&
		insertResp.Data.AccountID == insertReq.AccountID &&
		insertResp.Data.From == insertReq.From &&
		insertResp.Data.Consignee == insertReq.Consignee &&
		insertResp.Data.OrderID == insertReq.OrderID &&
		insertResp.Data.Phone == insertReq.Phone &&
		insertResp.Data.TargetDate == insertReq.TargetDate &&
		insertResp.Data.HandleDate == insertReq.HandleDate &&
		insertResp.Data.To == insertReq.To &&
		insertResp.Data.Weight == insertReq.Weight {
		isMatch = true
	}
	if !isMatch {
		log.Fatalf("Creation not match. Expect: %v, but get: %v", insertReq, insertResp.Data)
		return nil, err
	}
	//log.Fatalf("InsertConsignRecord response: %+v", insertResp)
	//existedConsign := insertResp.Data

	ctx.Set(ID, insertResp.Data.ID)
	ctx.Set(OrderID, insertResp.Data.OrderID)
	ctx.Set(AccountID, insertResp.Data.AccountID)
	ctx.Set(HandleDate, insertResp.Data.HandleDate)
	ctx.Set(TargetDate, insertResp.Data.TargetDate)
	ctx.Set(From, insertResp.Data.From)
	ctx.Set(To, insertResp.Data.To)
	ctx.Set(Consignee, insertResp.Data.Consignee)
	ctx.Set(Phone, insertResp.Data.Phone)
	ctx.Set(Weight, insertResp.Data.Weight)
	ctx.Set(IsWithin, insertResp.Data.IsWithin)

	return nil, nil
}

func QueryConsignPric(ctx *Context) (*NodeResult, error) {
	_, ok := ctx.Get(Client).(*service.SvcImpl)
	//cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part

	return nil, nil
}

func CreateConsignPrice(ctx *Context) (*NodeResult, error) {
	_, ok := ctx.Get(Client).(*service.SvcImpl)
	//cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part

	return nil, nil
}

// FoodBehaviorChain
func QueryFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Query all
	allFoodOrders, err := cli.FindAllFoodOrder()
	if err != nil {
		log.Fatalf("FindAllFoodOrder request failed, err %s", err)
		return nil, err
	}
	if len(allFoodOrders.Data) == 0 {
		log.Fatalf("FindAllFoodOrder returned empty results")
		return nil, err
	}
	if allFoodOrders.Status != 1 {
		log.Fatalf("FindAllFoodOrder failed: %v", allFoodOrders.Status)
		return nil, err
	}

	randomIndex := rand.Intn(len(allFoodOrders.Data))
	ctx.Set(OrderId, allFoodOrders.Data[randomIndex].OrderId)
	ctx.Set(FoodType, allFoodOrders.Data[randomIndex].FoodType)
	ctx.Set(StationName, allFoodOrders.Data[randomIndex].StationName)
	ctx.Set(StoreName, allFoodOrders.Data[randomIndex].StoreName)
	ctx.Set(FoodName, allFoodOrders.Data[randomIndex].FoodName)
	ctx.Set(Price, allFoodOrders.Data[randomIndex].Price)

	return nil, nil
}

func CreateFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock data
	MockedOrderID := ctx.Get(OrderId).(string)
	MockedID := faker.UUIDHyphenated()
	foodOrder := service.FoodOrder{
		ID:          MockedID,
		OrderID:     MockedOrderID,
		FoodType:    rand.Intn(1),
		FoodName:    generateRandomFood(),
		StationName: ctx.Get(StationName).(string),
		StoreName:   ctx.Get(StoreName).(string),
		Price:       ctx.Get(Price).(float64),
	}

	// Create Test
	newCreateResp, err := cli.CreateFoodOrder(&foodOrder)
	if err != nil {
		log.Fatalf("NewCreateFoodOrder request failed, err %s", err)
		return nil, err
	}
	if newCreateResp.Status != 1 {
		log.Fatalf("NEwCreateFoodOrder failed")
		return nil, err
	}

	ctx.Set(OrderId, newCreateResp.Data.OrderId)
	ctx.Set(FoodType, newCreateResp.Data.FoodType)
	ctx.Set(StationName, newCreateResp.Data.StationName)
	ctx.Set(StoreName, newCreateResp.Data.StoreName)
	ctx.Set(FoodName, newCreateResp.Data.FoodName)
	ctx.Set(Price, newCreateResp.Data.Price)

	return nil, nil
}

func QueryStationFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.GetAllStationFood()
	if err != nil {
		log.Fatalf("Resp returns err: %v", err)
		return nil, err
	}
	if resp.Status != 1 {
		log.Fatalf("GetAllStationFood status should be 1, but is %d", resp.Status)
		return nil, err
	}

	//Id           string  `json:"id"`
	//StationName  string  `json:"stationName"`
	//StoreName    string  `json:"storeName"`
	//Telephone    string  `json:"telephone"`
	//BusinessTime string  `json:"businessTime"`
	//DeliveryFee  float64 `json:"deliveryFee"`
	//FoodList     []struct {
	//	FoodName string  `json:"foodName"`
	//	Price    float64 `json:"price"`
	//} `json:"foodList"`

	// -------------------------------------------------------------- 01点31分

	return nil, nil
}

func QueryTrainFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part; I will generate it.

	return nil, nil
}

func QueryTrip(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	QueryAllTripResp, err := cli.QueryAllTrip()
	if err != nil {
		log.Fatalf("Request failed, err %s", err)
		return nil, err
	}
	if QueryAllTripResp.Status != 1 {
		log.Fatalf("Request failed, status: %d", QueryAllTripResp.Status)
		return nil, err
	}

	randomIndex := rand.Intn(len(QueryAllTripResp.Data))
	ctx.Set(TripID, QueryAllTripResp.Data[randomIndex].TripId)
	ctx.Set(From, QueryAllTripResp.Data[randomIndex].StartStationName)
	ctx.Set(From, QueryAllTripResp.Data[randomIndex].TerminalStationName)
	ctx.Set(Date, QueryAllTripResp.Data[randomIndex].StartTime)
	ctx.Set(StationName, QueryAllTripResp.Data[randomIndex].StationsName)
	ctx.Set(HandleDate, QueryAllTripResp.Data[randomIndex].EndTime)

	return nil, nil
}

func CreateTrip(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock para
	MockedLoginId := ctx.Get(LoginToken).(string)
	MockedTripId := GenerateTripId()
	MockedTrainTypeName := generateTrainTypeName(MockedTripId) /*"GaoTieSeven"*/
	MockedRouteID := ctx.Get(RouteID).(string)
	MockedStartStationName := ctx.Get(From).(string)
	MockedStationsName := /*strings.Join(AllRoutesByQuery.Data[0].Stations, ",")*/ ctx.Get(StationName).(string)
	MockedTerminalStationName := ctx.Get(To).(string)
	MockedStartTime := getRandomTime()
	MockedEndTime := getRandomTime(WithStartTime(MockedStartTime))

	// Mock input
	travelInfo := service.TravelInfo{
		LoginID:          MockedLoginId,
		TripID:           MockedTripId,
		TrainTypeName:    MockedTrainTypeName,
		RouteID:          MockedRouteID,
		StartStationName: MockedStartStationName,
		//StationsName:        fmt.Sprintf("%v,%v", MockedStartStationName, MockedTerminalStationName),
		StationsName:        MockedStationsName,
		TerminalStationName: MockedTerminalStationName,
		StartTime:           MockedStartTime,
		EndTime:             MockedEndTime,
	}

	// Create Test
	createResp, err := cli.CreateTrip(&travelInfo)
	if err != nil {
		log.Fatalf("CreateTrip request failed, err %s", err)
		return nil, err
	}
	if createResp.Status != 1 {
		log.Fatalf("CreateTrip failed: %s", createResp.Msg)
		return nil, err
	}
	if createResp.Msg == "Already exists" {
		log.Fatalf("Already exists: %s", createResp.Msg)
		return nil, err
	}
	isMatch := false
	if /*createResp.Data.Id == travelInfo.LoginID &&*/
	createResp.Data.StationsName == toLowerCaseAndRemoveSpaces(travelInfo.StationsName) &&
		createResp.Data.StartStationName == toLowerCaseAndRemoveSpaces(travelInfo.StartStationName) &&
		createResp.Data.TerminalStationName == toLowerCaseAndRemoveSpaces(travelInfo.TerminalStationName) &&
		createResp.Data.StartTime == travelInfo.StartTime &&
		createResp.Data.EndTime == travelInfo.EndTime &&
		createResp.Data.TrainTypeName == travelInfo.TrainTypeName &&
		createResp.Data.RouteId == travelInfo.RouteID {
		isMatch = true
	}
	if !isMatch {
		log.Fatalf("CreateTrip failed: %s. Except: %v, but get: %v", createResp.Msg, travelInfo, createResp.Data)
		return nil, err
	}

	ctx.Set(TripID, createResp.Data.TripId)
	ctx.Set(From, createResp.Data.StartStationName)
	ctx.Set(From, createResp.Data.TerminalStationName)
	ctx.Set(Date, createResp.Data.StartTime)
	ctx.Set(StationName, createResp.Data.StationsName)
	ctx.Set(HandleDate, createResp.Data.EndTime)

	return nil, nil
}

// TravelBehaviorChain
func QueryTrain(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part; I will generate it.

	return nil, nil
}

func QueryRoute(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	AllRoutesByQuery, err := cli.QueryAllRoutes()
	if err != nil {
		log.Fatalf("Request failed, err2 %s", err)
		return nil, err
	}
	if AllRoutesByQuery.Status != 1 {
		log.Fatal("AllRoutes_By_Query.Status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(AllRoutesByQuery.Data))
	ctx.Set(From, AllRoutesByQuery.Data[randomIndex].StartStation)
	ctx.Set(To, AllRoutesByQuery.Data[randomIndex].EndStation)
	ctx.Set(StationName, getMiddleElements(strings.Join(AllRoutesByQuery.Data[randomIndex].Stations, ",")))
	ctx.Set(RouteID, AllRoutesByQuery.Data[randomIndex].Id)

	return nil, nil
}

//func CreateRoute(ctx *Context) (*NodeResult, error) {
//	cli, ok := ctx.Get(Client).(*service.SvcImpl)
//	if !ok {
//		return nil, fmt.Errorf("service client not found in context")
//	}
//
//	// Create
//	MockedID := faker.UUIDHyphenated()
//	MockedStartStation := faker.GetRealAddress().City
//	MockedEndStation := faker.GetRealAddress().City
//	MockedStationList := fmt.Sprintf("%s,%s,%s", MockedStartStation, faker.GetRealAddress().City, MockedEndStation)
//	MockedDistanceList := fmt.Sprintf("%d,%d,%d", rand.Intn(30), rand.Intn(30), rand.Intn(30))
//	input := service.RouteInfo{
//		ID:           MockedID,
//		StartStation: MockedStartStation,
//		EndStation:   MockedEndStation,
//		StationList:  MockedStationList,
//		DistanceList: MockedDistanceList,
//	}
//	resp, err := cli.CreateAndModifyRoute(&input)
//	if err != nil {
//		log.Fatalf("Request failed, err %s", err)
//		return nil, err
//	}
//	if resp.Msg == "Already exists" {
//		log.Fatalf("Route already exists, skip")
//		return nil, err
//	}
//	if resp.Data.Id != input.ID {
//		log.Fatalf("Route ID does not match, expect %s, got %s", input.ID, resp.Data.Id)
//		return nil, err
//	}
//	if resp.Data.StartStation != input.StartStation {
//		log.Fatalf("StartStation does not match, expect %s, got %s", input.StartStation, resp.Data.StartStation)
//		return nil, err
//	}
//	if resp.Data.EndStation != input.EndStation {
//		log.Fatalf("StartStation does not match, expect %s, got %s", input.StartStation, resp.Data.StartStation)
//		return nil, err
//	}
//	if StringSliceToString(resp.Data.Stations) != ConvertCommaSeparatedToBracketed(input.StationList) {
//		log.Fatalf("StationList does not match, expect %s, got %s", ConvertCommaSeparatedToBracketed(input.StationList), StringSliceToString(resp.Data.Stations))
//		return nil, err
//	}
//	if IntSliceToString(resp.Data.Distances) != ConvertCommaSeparatedToBracketed(input.DistanceList) {
//		log.Fatalf("DistanceList does not match, expect %s, got %s", ConvertCommaSeparatedToBracketed(input.DistanceList), IntSliceToString(resp.Data.Distances))
//		return nil, err
//	}
//
//	ctx.Set(From, resp.Data.StartStation)
//	ctx.Set(To, resp.Data.EndStation)
//	ctx.Set(StationName, getMiddleElements(strings.Join(resp.Data.Stations, ",")))
//	ctx.Set(RouteID, resp.Data.Id)
//
//	return nil, nil
//}

func QueryBasic(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part; I will generate it.

	return nil, nil
}

func QuerySeat(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part; I will generate it.

	return nil, nil
}

// BasicBehaviorChain
func QueryStation(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part; I will generate it.

	return nil, nil
}

func QueryPrice(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part; I will generate it.

	return nil, nil
}

// SeatBehaviorChain
func QueryConfig(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part; I will generate it.

	return nil, nil
}

func QueryOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part; I will generate it.

	return nil, nil
}

func QueryOrderOther(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part; I will generate it.

	return nil, nil
}

/////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////

// Preserve Behaviors - The Last One
func Preserve(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	OrderTicketsInfo := service.OrderTicketsInfo{
		AccountID:       ctx.Get(AccountID).(string),  // Query:Create = 1 : 0
		ContactsID:      ctx.Get(ContactsID).(string), // Query:Create = 1 : 0
		TripID:          ctx.Get(TripID).(string),
		SeatType:        ctx.Get(SeatType).(int),
		LoginToken:      ctx.Get(LoginToken).(string),
		Date:            ctx.Get(Date).(string),
		From:            ctx.Get(From).(string),
		To:              ctx.Get(To).(string),
		Assurance:       ctx.Get(Assurance).(int),
		FoodType:        ctx.Get(FoodType).(int),
		StationName:     ctx.Get(StationName).(string),
		StoreName:       ctx.Get(StoreName).(string),
		FoodName:        ctx.Get(FoodName).(string),
		FoodPrice:       ctx.Get(FoodPrice).(float64),
		HandleDate:      ctx.Get(HandleDate).(string),
		ConsigneeName:   ctx.Get(ConsigneeName).(string),
		ConsigneePhone:  ctx.Get(ConsigneePhone).(string),
		ConsigneeWeight: ctx.Get(ConsigneeWeight).(float64),
		IsWithin:        ctx.Get(IsWithin).(bool),
	}
	PreserveResp, err := cli.Preserve(&OrderTicketsInfo)
	if err != nil {
		return nil, err
	}
	if PreserveResp.Status != 1 {
		return nil, fmt.Errorf("preserve order tickets fail. PreserveResp.Status != 1, get %v", PreserveResp.Status)
	}
	fmt.Printf("The Status is: %v, and PreserveResp Data: %v\n", PreserveResp.Status, PreserveResp.Data)
	fmt.Printf("PreserveBehaviors(Chain) Ends. End time: %v", time.Now().String())

	//return nil, nil
	return &(NodeResult{false}), nil // Chain End :D
}
