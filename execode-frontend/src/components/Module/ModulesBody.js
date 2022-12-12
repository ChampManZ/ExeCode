import * as React from 'react';
import { styled } from '@mui/material/styles';
import PropTypes from 'prop-types';
import Box from '@mui/material/Box';
import Collapse from '@mui/material/Collapse';
import IconButton from '@mui/material/IconButton';
import {Table, TableBody, TableCell, TableContainer, TableHead, TableRow} from '@mui/material';
import Paper from '@mui/material/Paper';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import modules from './AllModule';
import lecture1 from '../file/Lecture-01.pdf'

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
                                {module.name}
                            </TableCell> 
                            </>
                        </TableRow>

                        <TableRow>
                            <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
                                <Collapse in={open} timeout="auto" unmountOnExit>
                                    <Box sx={{ margin: 1 }}>
                                        <Table size="small" aria-label="lectures">
                                            <TableBody>
                                                {/* {modules.map(function(lecture){
                                                    return(
                                                        <TableRow key={lecture.name}> {lecture.detail}</TableRow>  
                                                    )
                                                })} */}
                                                <a href={lecture1} target='_blank' rel='noreferrer'></a>
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