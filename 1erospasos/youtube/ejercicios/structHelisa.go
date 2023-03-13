package ejercicios

type InsersionPedidos struct {
	Order Order `json:"order"`
}

type Order struct {
	AccountingDocument AccountingDocument `json:"accountingDocument"`
	Date               Date               `json:"date"`              
	QuestionGlobal     []QuestionGlobal   `json:"questionGlobal"`    
	ClientIdentity     string             `json:"clientIdentity"`    
	ClientCode         string             `json:"clientCode"`        
	IDVendor           string             `json:"idVendor"`          
	Transactions       []Transaction      `json:"transactions"`      
}

type AccountingDocument struct {
	Type   string `json:"type"`  
	Number string `json:"number"`
}

type Date struct {
	Day   string `json:"day"`  
	Month string `json:"month"`
	Year  string `json:"year"` 
}

type QuestionGlobal struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`  
}

type Transaction struct {
	Question    string `json:"question"`   
	Answer      string `json:"answer"`     
	Kind        int64  `json:"kind"`       
	Quantity    int64  `json:"quantity"`   
	IDProduct   int64  `json:"idProduct"`  
	Measure     int64  `json:"measure"`    
	IvaRate     string `json:"ivaRate"`    
	IvaValue    string `json:"ivaValue"`   
	UnitValue   int64  `json:"unitValue"`  
	TotalValue  int64  `json:"totalValue"` 
	PriceList   int64  `json:"priceList"`  
	RawValue    int64  `json:"rawValue"`   
	IDWarehouse int64  `json:"idWarehouse"`
}

// -------------
type InsersionTerceros struct {
	ThirdParty ThirdParty `json:"thirdParty"`
}

type ThirdParty struct {
	SecondLastName   string       `json:"secondLastName"`  
	Cree             bool         `json:"cree"`            
	FirstName        string       `json:"firstName"`       
	EconomicActivity string       `json:"economicActivity"`
	City             string       `json:"city"`            
	DocumentType     DocumentType `json:"documentType"`    
	Year             int64        `json:"year"`            
	Nature           Nature       `json:"nature"`          
	FirstLastName    string       `json:"firstLastName"`   
	Name             string       `json:"name"`            
	ID               string       `json:"id"`              
	CheckDigit       string       `json:"checkDigit"`      
	Regimen          DocumentType `json:"regimen"`         
	SecondName       string       `json:"secondName"`      
	UserCode         int64        `json:"userCode"`        
	Telephones       string       `json:"telephones"`      
	Cellphone        string       `json:"cellphone"`       
	Email            string       `json:"email"`           
	Address          string       `json:"address"`         
}

type DocumentType struct {
	Tag string `json:"tag"`
}

type Nature struct {
	Value int64 `json:"value"`
}

// --------------


type InsersionCliente struct {
	Client Client `json:"client"`
}

type Client struct {
	Portfolio            int64        `json:"portfolio"`           
	ID                   string       `json:"id"`                  
	CheckDigit           string       `json:"checkDigit"`          
	DocumentType         DocumentType `json:"documentType"`        
	Name                 string       `json:"name"`                
	Address              string       `json:"address"`             
	Telephones           string       `json:"telephones"`          
	Cellphone            string       `json:"cellphone"`           
	Fax                  string       `json:"fax"`                 
	City                 string       `json:"city"`                
	Email                string       `json:"email"`               
	Zone                 string       `json:"zone"`                
	Group                string       `json:"group"`               
	Free1                string       `json:"free1"`               
	Free2                string       `json:"free2"`               
	ToleranceDays        int64        `json:"toleranceDays"`       
	MaxCredit            int64        `json:"maxCredit"`           
	Regimen              DocumentType `json:"regimen"`             
	PaymentForm          int64        `json:"paymentForm"`         
	Cree                 bool         `json:"cree"`                
	EconomicActivity     int64        `json:"economicActivity"`    
	Nature               Nature       `json:"nature"`              
	FirstName            string       `json:"firstName"`           
	SecondName           string       `json:"secondName"`          
	FirstLastName        string       `json:"firstLastName"`       
	SecondLastName       string       `json:"secondLastName"`      
	Year                 int64        `json:"year"`                
	ApplyRTEFte          string       `json:"applyRteFte"`         
	ApplyMaxRTEFte       string       `json:"applyMaxRteFte"`      
	ApplyMaxRTEIva       string       `json:"applyMaxRteIva"`      
	ApplyMaxRTEIca       string       `json:"applyMaxRteIca"`      
	FiscalResponsibility int64        `json:"FiscalResponsibility"`
	FiscalRegime         int64        `json:"FiscalRegime"`        
	ContactName          string       `json:"contactName"`         
	ContactPhone         string       `json:"contactPhone"`        
	ContactEmail         string       `json:"contactEmail"`        
	AccountPortfolio     string       `json:"accountPortfolio"`    
	IDVendor             string       `json:"idVendor"`            
	PriceListServices    int64        `json:"priceListServices"`   
	PriceListArticules   int64        `json:"priceListArticules"`  
}

// type DocumentType struct {
// 	Tag string `json:"tag"`
// }

// type Nature struct {
// 	Value int64 `json:"value"`
// }

// ------