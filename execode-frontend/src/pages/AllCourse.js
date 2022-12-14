import CourseList from "../components/Course/CourseList"
import Sidebar from "../components/Sidebar"
import Box from '@mui/material/Box';
import { Outlet } from 'react-router-dom'

// import { ThemeProvider, useTheme } from "@emotion/react";
// import theme from "../components/theme";

function Courses () {
    return (
        <>
            <Sidebar></Sidebar>
            <Box sx={{ margin:15}}>
                <h1>Courses</h1>
                <CourseList></CourseList>
            </Box>
            <Outlet/>
        </>
        
    )
}

export default Courses