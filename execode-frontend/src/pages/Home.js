import { Box } from "@mui/material"
import Sidebar from "../components/Sidebar"
import TodoList from "../components/Home/TodoList"
import Announce from "../components/Home/Announce"
import axios from 'axios';
import jwt_decode from "jwt-decode";

// import { ThemeProvider } from "@emotion/react"
// import theme from "../components/theme"
// จะใช้ค่อย uncomment เอานะ

// Authenticate Zone
var userdata = {}

// Get code to gain token
const queryParameters = new URLSearchParams(window.location.search);
const code = queryParameters.get("code");

const paramsObj = { grant_type: "authorization_code", client_id: "5ujg8c9bbiihttb1gdsijj48t0", redirect_uri: "http://localhost:3001/home", code: code };
const params = new URLSearchParams(paramsObj);
const headerAdd = {
    headers: {
        'Authorization': `Basic NXVqZzhjOWJiaWlodHRiMWdkc2lqajQ4dDA6ZmY4ZGc5M3ZwOGpha2o5YmVhdms2NGY5MWZvOXFjMzhsZjg2dXV1cWhvczFwOGNmOHVo`
    }
}

axios.post('https://execode-users.auth.us-east-1.amazoncognito.com/oauth2/token', params, headerAdd)
.then(function(res) {
    extractData(res.data.id_token)
    postUser(res.data.access_token)
})
.catch(err => console.log('Failed to get token'));

// Extract Data
function extractData(myToken) {
    var token = myToken;
    var decoded = jwt_decode(token);
    var dec_name = 'cognito:username'
    userdata.user_name = decoded[dec_name];
    userdata.name = decoded.name
    userdata.email = decoded.email
}

// Post data
const postUser = function (token) {
    axios.post('http://localhost:3000/users', userdata, {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    }).then(res => console.log(res)).catch(err => console.log(err));
}

function Home () {
    return (
        <>
        <Sidebar />
        <Box sx={{ margin:15 }}>
            <h1>To Do...</h1>
            <TodoList />
            <Box sx={{ mt:10 }}>
            <h1>Announcement</h1>
            <Announce />
            </Box>
        </Box>
        </>
        
    )
}


export default Home