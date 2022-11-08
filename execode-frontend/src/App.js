import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Coderunner from "./pages/Coderunner";
import Courses from "./pages/Courses";
import Home from "./pages/Home";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='courses' element={<Courses />} />
          <Route path='/coderunplayground' element={<Coderunner />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
