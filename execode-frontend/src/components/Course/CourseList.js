import * as React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import {Grid} from '@mui/material'
import { CardActionArea } from '@mui/material';
import courses from './Courses';

export default function Course() {
  return (
    <>
    <Grid container columnSpacing={5} rowSpacing={5}>
      {courses.map(function(course) {
        return (
          <Grid item sm={6} md={4} lg={3}>
            <Card sx={{ maxWidth: 345, minHeight: 330 }}>
              <CardActionArea href='/courses/1'>
                <CardMedia
                  component="img"
                  height="140"
                  image="/image/course_1.jpg"
                  alt="course_1" />
                  
                <CardContent>
                  <>
                  <Typography gutterBottom variant="h5" component="div">
                    {course.name}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    {course.description}
                  </Typography>
                  </>
                </CardContent>
              </CardActionArea>
            </Card>    
          </Grid>
        )
      })}
    </Grid>
      {/* <Grid item sm={6} md={4} lg={3}>
        <Card sx={{ maxWidth: 345, minHeight: 330 }}>
          <CardActionArea href='/courses/2'>
            <CardMedia
              component="img"
              height="140"
              image="/image/course_1.jpg"
              alt="course_1" />
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                Basic Python
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Lizards are a widespread group of squamate reptiles, with over 6,000
                species, ranging across all continents except Antarctica
              </Typography>
            </CardContent>
          </CardActionArea>
        </Card>
      </Grid>

      <Grid item sm={6} md={4} lg={3}>
        <Card sx={{ maxWidth: 345, minHeight: 330 }}>
          <CardActionArea>
            <CardMedia
              component="img"
              height="140"
              image="/image/course_1.jpg"
              alt="course_1" />
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                Fundamental of Programming
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Lizards are a widespread group of squamate reptiles, with over 6,000
                species, ranging across all continents except Antarctica
              </Typography>
            </CardContent>
          </CardActionArea>
        </Card>
      </Grid>

      <Grid item sm={6} md={4} lg={3}>
        <Card sx={{ maxWidth: 345, minHeight: 330 }}>
          <CardActionArea>
            <CardMedia
              component="img"
              height="140"
              image="/image/course_1.jpg"
              alt="course_1" />
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                Fundamental of Programming
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Lizards are a widespread group of squamate reptiles, with over 6,000
                species, ranging across all continents except Antarctica
              </Typography>
            </CardContent>
          </CardActionArea>
        </Card>
      </Grid>

      <Grid item sm={6} md={4} lg={3}>
        <Card sx={{ maxWidth: 345, minHeight: 330 }}>
          <CardActionArea>
            <CardMedia
              component="img"
              height="140"
              image="/image/course_1.jpg"
              alt="course_1" />
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                Fundamental of Programming
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Lizards are a widespread group of squamate reptiles, with over 6,000
                species, ranging across all continents except Antarctica
              </Typography>
            </CardContent>
          </CardActionArea>
        </Card>
      </Grid> */}

    </>
  );
}



