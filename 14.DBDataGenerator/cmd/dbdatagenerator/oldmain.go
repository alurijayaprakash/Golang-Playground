package main

// Note : to insert 1 lakh DB records into DB, I tried to implement Concurrency in it.
// but infact SQL inserts are purely i/o operations but not CPU intensive tasks.

// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"sync"
// 	"time"

// 	"github.com/brianvoe/gofakeit/v6"
// 	_ "github.com/lib/pq"
// )

// // Student struct represents the structure of the data you want to insert
// type Student struct {
// 	StudentID int    `json:"studentid"`
// 	FirstName string `json:"firstname"`
// 	Email     string `json:"email"`
// 	RollNo    int    `json:"rollno"`
// 	PanNo     string `json:"panno"`
// 	SSNID     string `json:"ssnid"`
// 	Gender    string `json:"gender"`
// 	Profile   json.RawMessage
// }

// type StudentRaw struct {
// 	StudentID int    `json:"studentid"`
// 	FirstName string `json:"firstname"`
// 	Email     string `json:"email"`
// 	RollNo    int    `json:"rollno"`
// 	PanNo     string `json:"panno"`
// 	SSNID     string `json:"ssnid"`
// 	Gender    string `json:"gender"`
// }

// // generateRandomPAN generates a random PAN number
// func generateRandomPAN() string {
// 	// PAN format: ABCDE1234F
// 	// A to Z: 65 to 90 (ASCII values)
// 	// 0 to 9: 48 to 57 (ASCII values)

// 	rand.Seed(time.Now().UnixNano())

// 	// Generate the first five characters (alphabets)
// 	firstFive := ""
// 	for i := 0; i < 5; i++ {
// 		randomChar := rune(rand.Intn(26) + 65) // ASCII values for A to Z
// 		firstFive += string(randomChar)
// 	}

// 	// Generate the next four characters (digits)
// 	nextFour := fmt.Sprintf("%04d", rand.Intn(10000)) // Random 4-digit number

// 	// Generate the last character (alphabet)
// 	lastChar := rune(rand.Intn(26) + 65) // ASCII values for A to Z

// 	// Concatenate the parts to form the PAN number
// 	panNumber := fmt.Sprintf("%s%s%s", firstFive, nextFour, string(lastChar))

// 	return panNumber
// }

// // generate random num
// func generateRandomNumber(min, max int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(max-min+1) + min
// }

// // generateRandomGender generates a random gender ("M" or "F")
// func generateRandomGender() string {
// 	rand.Seed(time.Now().UnixNano())
// 	genders := []string{"M", "F"}
// 	return genders[rand.Intn(len(genders))]
// }

// type Datadb struct {
// 	DB *sql.DB
// }

// var wg = sync.WaitGroup{}

// func main() {
// 	DBRecordsCount := 100000
// 	ddb := new(Datadb)
// 	// Connection parameters
// 	connStr := "user=postgres password=root dbname=College sslmode=disable"

// 	// Open a database connection
// 	DB, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ddb.DB = DB

// 	// Example of setting a maximum connection pool size
// 	// ddb.DB.SetMaxOpenConns(10)

// 	defer ddb.DB.Close()

// 	// current time
// 	startTime := time.Now()
// 	wg.Add(DBRecordsCount)
// 	for i := 0; i < DBRecordsCount; i++ {
// 		fmt.Println("PushData2DB :", i)
// 		ddb.PushData2DB(i)
// 		// fmt.Println(i)
// 	}
// 	// Wait for all goroutines to finish
// 	wg.Wait()

// 	elapsedTime := time.Since(startTime)

// 	fmt.Printf("Time taken: %s\n", elapsedTime)

// }

// // Insert Func
// func (d *Datadb) PushData2DB(i int) {

// 	MyStudentID := generateRandomNumber(1, 99999)
// 	Myfirstname := gofakeit.Name() // Markus Moen
// 	Myemail := gofakeit.Email()    // alaynawuckert@kozey.biz
// 	MySsn := gofakeit.SSN()        // 578-23-4577
// 	MypanNumber := generateRandomPAN()
// 	MyRollNum := generateRandomNumber(100, 9999)
// 	MyGender := generateRandomGender()
// 	studentraw := StudentRaw{
// 		StudentID: MyStudentID,
// 		FirstName: Myfirstname,
// 		Email:     Myemail,
// 		RollNo:    MyRollNum,
// 		PanNo:     MypanNumber,
// 		SSNID:     MySsn,
// 		Gender:    MyGender,
// 	}

// 	// Convert Student object to JSON
// 	studentrawJSON, err := json.Marshal(studentraw)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON:", err)
// 		return
// 	}

// 	// Convert JSON to RawMessage
// 	MyRawMessage := json.RawMessage(studentrawJSON)

// 	// Insert data into the student table in the "iiith" schema
// 	student := Student{
// 		StudentID: MyStudentID,
// 		FirstName: Myfirstname,
// 		Email:     Myemail,
// 		RollNo:    MyRollNum,
// 		PanNo:     MypanNumber,
// 		SSNID:     MySsn,
// 		Gender:    MyGender,
// 		Profile:   MyRawMessage,
// 	}

// 	err = d.insertStudent(student)
// 	if err != nil {
// 		fmt.Println("ERROR")
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Record inserted successfully! : ", i)
// 	wg.Done()
// }

// // insertStudent inserts a Student record into the student table in the "iiith" schema
// func (d *Datadb) insertStudent(s Student) error {
// 	query := `
// 		INSERT INTO iiith.studentnew (StudentID, FirstName, Email, RollNo, PanNo, SSNID, Gender, profile)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
// 	`

// 	// Marshal JSON data for the profile column
// 	profileJSON, err := json.Marshal(s.Profile)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = d.DB.Exec(query, s.StudentID, s.FirstName, s.Email, s.RollNo, s.PanNo, s.SSNID, s.Gender, profileJSON)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// DDL Postgresql
// CREATE TABLE studentnew (
// 	StudentID INT,
// 	FirstName VARCHAR(255),
// 	Email VARCHAR(255),
// 	RollNo INT,
// 	PanNo VARCHAR(255),
// 	SSNID VARCHAR(255),
// 	Gender VARCHAR(255),
// 	profile jsonb NULL
// 	);
