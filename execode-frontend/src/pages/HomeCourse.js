import Box from '@mui/material/Box';
import { Stack } from '@mui/system';
import { Typography } from '@mui/material';
import { Outlet } from 'react-router-dom';
import Sidebar from '../components/Sidebar';
import CourseMenu from '../components/Course/CourseMenu';

export default function HomeCourse () {

    return (
        <>
            <Sidebar></Sidebar>
            <Stack direction="row" spacing={2} justifyContent="space-between" sx={{mt:15}}>
                <CourseMenu/>
                <Box flex={5} sx={{mt:15}}>
                    <Typography color="text.primary">
                        <h1>Fundamental of Programming</h1>
                    </Typography>

                </Box>
            </Stack>
            <Outlet/>
        </>
        
    
    )
}