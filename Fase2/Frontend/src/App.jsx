import { useState } from 'react';
import {Route, Routes} from 'react-router-dom';

import Estudiante from './pages/Estudiante';
import Login from './pages/login';

//import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
    <Routes>
    <Route path='/' element={<Login/>}/>
      <Route path='/principal/estudiante' element={<Estudiante/>}/>
    </Routes>
    </>

  )
}

export default App
