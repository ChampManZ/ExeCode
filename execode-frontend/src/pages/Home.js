import { Box } from "@mui/material"
import Sidebar from "../components/Sidebar"
import TodoList from "../components/Home/TodoList"
import Announce from "../components/Home/Announce"
// import { ThemeProvider } from "@emotion/react"
// import theme from "../components/theme"
// จะใช้ค่อย uncomment เอานะ



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