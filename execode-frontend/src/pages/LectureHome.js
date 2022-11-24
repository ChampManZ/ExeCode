import { Box } from "@mui/material"
import Sidebar from "../components/Sidebar"
import TodoList from "../components/Home/TodoList"
import LectureSchedule from "../components/Home/LectureSchedule"

function Home () {

    return (
        <>
        <Sidebar />
        <Box sx={{ margin:15 }}>
            <h1>To Do...</h1>
            <TodoList />
            <Box sx={{ mt:10 }}>
            <h1>Scheduled</h1>
            <LectureSchedule />
            </Box>
        </Box>
        </>
        
    )
}


export default Home