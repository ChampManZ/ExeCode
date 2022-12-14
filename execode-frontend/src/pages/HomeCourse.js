import Box from '@mui/material/Box';
import { Stack } from '@mui/system';
import { Divider, Typography } from '@mui/material';
import { Outlet } from 'react-router-dom';
import Sidebar from '../components/Sidebar';
import CourseMenu from '../components/Course/CourseMenu';
import courses from '../components/Course/Courses';
import '../styles/homecourse.css'

export default function HomeCourse () {

    return (
        <>
            <Sidebar></Sidebar>
            <Stack direction="row" spacing={2} justifyContent="space-between" sx={{mt:15, mr:30}}>
                <CourseMenu/>
                <Box flex={5} sx={{mt:15}}>
                    <Typography color="text.primary">
                        <h1>Fundamental of Programming</h1>
                    </Typography>
                    <Divider />
                    
                    <Box sx={{mt:10, mb:5}}>
                        <img className='homecourse-img' src='../image/program banner.jpg' alt="programming banner"></img>
                    </Box>

                    <Box sx={{mb:15}}>
                      {courses.map(function(course) {
                        return (
                            <p className='homecourse-p'>{course.about}</p>
                        )
                    })}   
                    </Box>
                    
                </Box>
            </Stack>
            <Outlet/>
        </>
        
    
    )

}