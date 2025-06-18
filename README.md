mySecondAwesomeProject
একটা ছোট Go প্রজেক্ট যেখানে আমরা আলাদা আলাদা প্যাকেজ ও মেথড ইউজ করে struct এবং method এর access control (public/private) শিখব।

Project Structure (প্রজেক্টের ফোল্ডার গঠন)
go
Copy
Edit
mySecondAwesomeProject/
├── main.go                 (main package)
├── person/
│   └── person.go           (person package)
└── new_package/
    └── new_package.go      (new_package package)
Project Description (প্রজেক্ট বর্ণনা)
এই প্রজেক্টে:

আমরা Person নামে একটা struct বানিয়েছি যেটা person প্যাকেজে আছে।

Person struct এর private method greet() এবং public method Meet() আছে।

new_package প্যাকেজ থেকে আমরা Try() ফাংশন কল করি যা Person struct ইউজ করে।

main.go থেকে আমরা এই দুই প্যাকেজের ফাংশন ও struct ব্যবহার করি।

এভাবে আমরা Go তে package, access control (public/private) এবং modular কোড লেখা শিখব।

কীভাবে রান করবেন (How to Run)
প্রথমে project folder এ terminal/command prompt ওপেন করুন।

Go module init করুন:

go
Copy
Edit
go mod init mySecondAwesomeProject
go mod tidy
প্রজেক্ট রান করুন:

go
Copy
Edit
go run main.go
Expected Output (আউটপুট)
pgsql
Copy
Edit
Main running
From public method, nice to meet you!
Hello from private method greet: Sabbir
Try function from new_package finished
Welcome to second phase 
From public method, nice to meet you!
Hello from private method greet: Sabbir Dulabhai
Important Notes (গুরুত্বপূর্ণ নোটস)
Go তে প্যাকেজের নাম আর ফাইল structure ঠিক রাখতে হবে।

এক প্যাকেজের private method অন্য প্যাকেজ থেকে access করা যায় না।

Public method গুলো capital letter দিয়ে শুরু করতে হয় যাতে অন্য প্যাকেজ থেকে access করা যায়।

go.mod ফাইল তৈরি করে module নাম ঠিক রাখতে হবে।

Contact (যোগাযোগ)
এই প্রজেক্টের জন্য সাহায্যের দরকার হলে যোগাযোগ করুন।
saidurnsu.dev@gmail.com

Thank you for using this project!

