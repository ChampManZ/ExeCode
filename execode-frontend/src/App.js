import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Coderunner from "./pages/Coderunner";
import AllCourse from "./pages/AllCourse";
import Home from "./pages/Home";
import ModuleCourse from "./pages/ModuleCourse";
import HomeCourse from './pages/HomeCourse';


function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/courses'>
            <Route index element={<AllCourse />} />
            <Route path=":id">
              <Route index element={<HomeCourse/>} />
              <Route path='module' element={<ModuleCourse/>} />
            </Route> 
          </Route>
          <Route path='/coderunplayground' element={<Coderunner />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
