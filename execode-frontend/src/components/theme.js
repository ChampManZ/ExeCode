import { createTheme } from "@mui/system";
<link
  rel="stylesheet"
  href="https://fonts.googleapis.com/css?family=Amiko"
/>

const theme = createTheme({
    typography: {
        h1: {
            fontFamily: 'Amiko',
            fontSize: "36px"
        }
    },

    palette:{
       primary: {
        main: '#8EDFC7',
        contrastText: '#000000'
       }
    }
    
})

export default theme