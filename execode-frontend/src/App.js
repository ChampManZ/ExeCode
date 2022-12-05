import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Coderunner from "./pages/Coderunner";
import AllCourse from "./pages/AllCourse";
import Home from "./pages/Home";
import ModuleCourse from "./pages/ModuleCourse";
import HomeCourse from './pages/HomeCourse';
import LectureHome from './pages/LectureHome';
import Welcome from './pages/Welcome';

function App() {

  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Welcome />}/>
          <Route path='/home' element={<Home />} />
          <Route path='/courses'>
            <Route index element={<AllCourse />} />
            <Route path=":id">
              <Route index element={<HomeCourse/>} />
              <Route path='module' element={<ModuleCourse/>} />
            </Route> 
          </Route>
          <Route path='/coderunplayground' element={<Coderunner />} />
          <Route path='/test1' element={<LectureHome />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
