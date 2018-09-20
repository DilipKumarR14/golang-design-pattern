package main
import (
	"encoding/json"
	"io/ioutil"
	"os"  
	"fmt"
)
type Address struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Mobile int `json:"mobile"`
	State string `json:"state"`
	City string `json:"City"`
	Streetno int `json:"Streetno"`
	Streetname string `json:"Streetname"`
	Zipcode int`json:"zip"`
} 

func Create()  {
	var book Address	
	
	fmt.Println("enter the first name")
	fmt.Scanf("%s",&book.Firstname)
	
	fmt.Println("enter the last name")
	fmt.Scanf("%s",&book.Lastname)
	fmt.Println("enter the mobile number")
	fmt.Scanf("%d",&book.Mobile)
	fmt.Println("enter the state name")
	fmt.Scanf("%s",&book.State)
	fmt.Println("enter the City name")
	fmt.Scanf("%s",&book.City)
	fmt.Println("enter the Streetno")
	fmt.Scanf("%d",&book.Streetno)
	fmt.Println("enter the Streetname")
	fmt.Scanf("%s",&book.Streetname)
	fmt.Println("enter the zip")
	fmt.Scanf("%d",&book.Zipcode)
	
	file, err1 := ioutil.ReadFile("file.json")
if err1 != nil {
os.Exit(1)
}
var arr []Address

err2 := json.Unmarshal(file, &arr)
if err2 != nil {
fmt.Println("error:", err2)
os.Exit(1)
}
arr = append(arr, Address{
	Firstname:book.Firstname,
	Lastname:book.Lastname,
	Mobile:book.Mobile,
	State:book.State,
	City:book.City,
	Streetno:book.Streetno,
	Streetname:book.Streetname,
	Zipcode:book.Zipcode,
	})

result, err := json.Marshal(arr)
if err != nil {
fmt.Println(err)
}
err = ioutil.WriteFile("file.json", result, 0644)

fmt.Printf("%+v", string(result))
for i := range arr {
fmt.Printf("firstName: %s\n", arr[i].Firstname)
fmt.Printf("lastName: %s\n", arr[i].Lastname)
}
	// buf:=new(bytes.Buffer)
	// encoder:=json.NewEncoder(buf)
	// encoder.Encode(b1)
	// file,err:=os.Create("file.json")
	// if(err!=nil){
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// io.Copy(file,buf) 

	// data,err:=json.Marshal(b1)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(data))

	// xp:=[]Address{b1}
	// fmt.Println("go data: %+v",xp)
	// data,err:=json.Marshal(b1)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println("json: ",string(data))
}

func Delete(){

}


  
func main(){
Create()
}