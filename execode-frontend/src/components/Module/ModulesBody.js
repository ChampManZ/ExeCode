import * as React from 'react';

import Box from '@mui/material/Box';
import Collapse from '@mui/material/Collapse';
import IconButton from '@mui/material/IconButton';
import {Table, TableBody, TableCell, TableContainer, TableRow} from '@mui/material';
import Paper from '@mui/material/Paper';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import modules from './AllModule';
import '../../styles/module.css'

export default function ModulesBody() {
    const [open, setOpen] = React.useState(false);

    return (
        <>
        {modules.map(function(module, i){
            return(
                <Box sx={{mb:5, width: '75%'}}>
                    <TableContainer component={Paper}>
                        <TableRow sx={{ '& > *': { borderBottom: 'unset' } }}>
                            <>
                            <TableCell>
                                <IconButton
                                    aria-label="expand row"
                                    size="small"
                                    onClick={() => setOpen(!open)}
                                >
                                    {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
                                </IconButton>
                            </TableCell>

                            <TableCell key={i} component="th" scope="row">
                                <h2 className='module-h2'>{module.name}</h2>
                            </TableCell> 
                            
                            </>
                        </TableRow>

                        <TableRow>
                            <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
                                <Collapse in={open} timeout="auto" unmountOnExit>
                                    <Box sx={{ margin: 1 }}>
                                        <Table size="small" aria-label="lectures">
                                            <TableBody>
                                                {modules.map(function(lecture, i){
                                                    return(
                                                        <TableRow key={i}> 
                                                            <a className='module-body' href={lecture.link} target='_blank' rel='noreferrer'>{lecture.detail}</a>
                                                            
                                                        </TableRow>  
                                                    )
                                                })}
                                            </TableBody> 
                                        </Table>
                                    
                                    </Box>  
                                </Collapse>
                            </TableCell>
                        </TableRow>
                        
                    </TableContainer>
                </Box>
            );
        })}
        </>
    );
}