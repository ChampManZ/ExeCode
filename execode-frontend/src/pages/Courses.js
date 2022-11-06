import CourseList from "../components/Course/CourseList"
import Sidebar from "../components/Sidebar"
import Box from '@mui/material/Box';
import { Typography, } from "@mui/material";
import { ThemeProvider, useTheme } from "@emotion/react";
import theme from "../components/theme";

function Courses () {
    return (
        <>
            <Sidebar></Sidebar>
            <Box sx={{ margin:15}}>
                <Typography variant="h1">Courses</Typography>
                <CourseList></CourseList>
            </Box>
        </>
        
    )
}

export default Courses