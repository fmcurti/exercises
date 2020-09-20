import React, { useState } from 'react'
import ColoredCircle from './Components/ColoredCircle.js'
import ColorForm from './Components/ColorForm.js'




const App = () => {

    const [color,setColor] = useState('black')

    return (
        <div>
        <ColoredCircle color={color}/>
        <h3>Type a Color!</h3> 
        <ColorForm value={color} setColor={setColor}/>
        </div>
    )
}


export default App