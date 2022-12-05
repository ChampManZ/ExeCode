import * as React from 'react';
import { styled } from '@mui/material/styles';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell, { tableCellClasses } from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';


const StyledTableCell = styled(TableCell)(({ theme }) => ({
  [`&.${tableCellClasses.head}`]: {
    backgroundColor: theme.palette.common.black,
    color: theme.palette.common.white,
  },
  [`&.${tableCellClasses.body}`]: {
    fontSize: 14,
  },
}));

const StyledTableRow = styled(TableRow)(({ theme }) => ({
  '&:nth-of-type(odd)': {
    backgroundColor: theme.palette.action.hover,
  },
  // hide last border
  '&:last-child td, &:last-child th': {
    border: 0,
  },
}));

function createData(course, task, duedate) {
  return { course, task, duedate };
}

const rows = [
  createData('Fundamental of Programming', 'Complete Module 6', Date()),
  createData('Cyber-Physical System', 'Complete Module 6', Date()),
  createData('Computer System', 'Complete Module 6', Date()),
  createData('Discrete Mathematics', 'Complete Module 6', Date()),
  createData('Database Technology', 'Complete Module 6', Date()),
];

export default function CustomizedTables(props) {
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 700 }} aria-label="customized table">
        <TableHead>
          <TableRow>
            <StyledTableCell>Course</StyledTableCell>
            <StyledTableCell align="left">Task</StyledTableCell>
            <StyledTableCell align="left">Due Date</StyledTableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <StyledTableRow key={row.course}>
              <StyledTableCell component="th" scope="row">
                {row.course}
              </StyledTableCell>
              <StyledTableCell align="left">{row.task}</StyledTableCell>
              <StyledTableCell align="left">{row.duedate}</StyledTableCell>
            </StyledTableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
