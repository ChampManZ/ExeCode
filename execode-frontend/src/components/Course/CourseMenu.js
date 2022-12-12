import { ListItem, ListItemButton, ListItemText, Toolbar } from '@mui/material';
// import { useParams } from "react-router-dom"
import List from '@mui/material/List';
import Box from '@mui/material/Box';

export default function CourseMenu() {
    // const {id} = useParams()
    return (
        <Box flex={1} sx={{mt:10}}>
            <Toolbar position="sticky">
                <List>
                    <ListItem>
                        <ListItemButton href="/courses">
                            <ListItemText primary="Course"/>
                        </ListItemButton>
                    </ListItem>

                    <ListItem>
                        <ListItemButton href="module">
                            <ListItemText primary="Modules"/>
                        </ListItemButton>
                    </ListItem>

                    {/* <ListItem>
                        <ListItemButton>
                            <ListItemText primary="Assignment"/>
                        </ListItemButton>
                    </ListItem> */}
                </List>
            </Toolbar>
        </Box>
    )
}