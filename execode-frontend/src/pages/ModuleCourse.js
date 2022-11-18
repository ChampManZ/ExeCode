import Sidebar from "../components/Sidebar"
import CourseMenu from "../components/Course/CourseMenu"
import HeadNav from "../components/Course/HeadNav"
import { Stack } from "@mui/material"


function ModuleCourse () {
    return (
        <>
            <Sidebar></Sidebar>
            <Stack direction="row" spacing={2} justifyContent="space-between" sx={{mt:15}}>
                <CourseMenu/>
                <HeadNav/>
            </Stack>

        </>
        )

}

export default ModuleCourse
