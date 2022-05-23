## Inspiration/Rationale
We believe that finding parking is becoming increasingly difficult, especially as more drivers get back on the road. On top of this, paying for parking can be a hassle - having to deal with parking meters that only take coins (who carries coins anymore?) or bulky parking machines that make you type in your car details, which can waste a lot of time.

We saw that in amusement parks, there are QR or barcode scanners that easily confirm if a guest has a valid reservation or ticket, which took 1-2 seconds to scan through. Surprisingly enough, the scanner machines were similarly shaped to parking meters too.

With that said, our goal with this hackathon was to come up with an easy-to-use solution to make peoples’ lives easier. The solution we envisioned came in the form of a mobile app, ParkIt. 

## What it does
ParkIt generates a unique QR code that drivers can use to scan in and out of a parking position, simplifying the parking process by introducing efficient organization. When a user signs in, they are assigned a QR code, which is stored in a database. 

Upon initial scan in, the QR code is registered as 'active' if a parking position can be assigned. If there are no positions available within that parking lot, the app will let the user know. In this way, the driver can save time by driving elsewhere instead of entering the parking lot and scouting for an open position. The user at this time may then proceed to park in the position indicated. 

When the user is leaving, a subsequent scan out marks the QR code as 'inactive', and the parking spot that was previously occupied is now marked vacant. 

## How we built it: Frontend
Our frontend was developed using the Flutter framework which utilizes the Dart language to create comprehensive mobile interfaces while maintaining strong performance and easy to understand syntax. The wide array of documentation online as well as the regular support this framework receives from Google makes it a very strong alternative for mobile application development.

One of our goals for this hackathon was to expose ourselves to new technologies. With that said, all members within our team attempted to utilize tools that we have little to no experience with. As a result, the front-end side of our mobile application has only been partially developed. Therefore, the majority of our demo will utilize the Figma prototype that we designed to display some of the core features of our application. 

Regardless, we would like to applaud Kin,  our Frontend developer, for their efforts in this hackathon as not only was it their first time coding in a full stack environment, but it was also their first time coding in general. It has been an amazing opportunity to revisit the struggles and triumphs of learning how to code a large scale project for the first time.

## How we built it: Backend
While we felt that the Flutter framework was needed to create an effective frontend for our project, we believed that our backend was small scale enough (at least in its current iteration) to forgo the need for a backend framework. As mentioned above, we sought to utilize tools in which we had no experience with. Golang is an open-source programming language that has been steadily increasing in popularity and was the selected choice for our backend. 

A local server and its' endpoints was created in Golang. This server would then be made public using NGROK, a software that exposes local ports to the Internet. In order to connect our server to a MySQL database for user authentication and QR code storage, we utilized a public Go-SQL driver (https://github.com/go-sql-driver/mysql). 

By leveraging the capabilities of the 'net/http' package in Golang, we could establish communication between our server and the frontend via the use of JSON request and response objects. 

## What's next for ParkIt! - FireTrucks™
For the time being, this project will remain as a simple component to a larger solution. We have not decided if we would continue to further develop it, but at the very least, this project is sufficient to act as a demo of working technology! 
