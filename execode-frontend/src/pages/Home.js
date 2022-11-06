import { Box } from "@mui/material"
import Sidebar from "../components/Sidebar"
import TodoList from "../components/Home/TodoList"
import { ThemeProvider } from "@emotion/react"
import theme from "../components/theme"



function Home () {

    return (
        <>
        <Sidebar></Sidebar>
        <Box sx={{ margin:15 }}>
            <h1>To Do...</h1>
            <TodoList></TodoList>

            <Box sx={{ mt:10 }}>
            <h1>Announcement</h1>
            </Box>
        </Box>
        </>
        
    )
}


export default Home