import Sidebar from "../components/Sidebar"
import CourseMenu from "../components/Course/CourseMenu"
import HeadNav from "../components/Module/HeadNav"
// import ModulesBody from "../components/Module/ModulesBody"
import '../styles/module.css'
import { Stack } from "@mui/material"


function ModuleCourse () {
    return (
        <>
            <Sidebar></Sidebar>
            <Stack direction="row" spacing={2}  sx={{mt:15}}>
                <CourseMenu/> 
                <HeadNav/>  
            </Stack>

        </>
        )

}

export default ModuleCourse
