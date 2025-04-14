Overview of the UI and backend : 
In my React application, I am using Leaflet as the core mapping library to display and interact with maps using OpenFreeMaps . (Some errors were occuring with OPenFREEmaps so I used OpenStreetMaps for the task). 
I have integrated frontend and backend in my code wherein for the frontend, I have created 3 pages for-Signup,Login,Maps. 
The signup page takes in the name,email,password of the user and stores the data in a PSQL database named 'evaluation' while the login page asks only for the 'email' and password and if the user is found to be authentic,the user is directed to the maps page and a jwt token is sent.
Though I have not yet done implementation of the part wherein features are restricted to logged in users. 

Regarding the maps :
The MapView component uses Leaflet's MapContainer, TileLayer, and Marker components (from react-leaflet) to: 
a) Render an interactive map (using OpenStreetMap tiles by default). 
b) Display markers at specific coordinates. 
c) Show popups when markers are clicked. 

User WOrkflow : User searches for a place → the react app fetches coordinates (through Nominatim API provided by OpenStreetMaps) 
Leaflet updates the map: Centers the view on those coordinates.

Folder Chronology : I have divided the main folder into 2 main subfolders : backend and maps. 
The backend subfolder has the backend for storing the data of users using Go and Gin. It contaings the main.go file. 
It has various subfolders like config(for database handling),utils(for jwt handling),etc.
The maps folder is the react application. It has src folder which has the App.jsx file . Additionaly, I have created a components folder which has various file/Component codes for rendering the UI. 

You are requested to navigate to the maps subfolder to start the UI application.

