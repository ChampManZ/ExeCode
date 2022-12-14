import * as React from 'react';
import Breadcrumbs from '@mui/material/Breadcrumbs';
import Typography from '@mui/material/Typography';
import Link from '@mui/material/Link';
import NavigateNextIcon from '@mui/icons-material/NavigateNext';
import Box from '@mui/material/Box';
import ModulesBody from './ModulesBody';




export default function HeadNav () {

    function handleClick(event) {
        event.preventDefault();
        console.info('You clicked a breadcrumb.');
      }
    
    const breadcrumbs = [
    <Link underline="hover" key="1" color="inherit" href="/course" onClick={handleClick}>
        <h1>Fundamental of Programming</h1>
    </Link>,
   
    <Typography key="2" color="text.primary">
        <h1>Modules</h1>
    </Typography>,
    ];

    return (
        <Box flex={5} sx={{mt:15}}>
            <Breadcrumbs
                separator={<NavigateNextIcon fontSize="medium" />}
                aria-label="breadcrumb"
                sx={{my:5}}
            >
                {breadcrumbs}
            </Breadcrumbs>
            <ModulesBody/>
        </Box>
    )
}