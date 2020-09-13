import React, { useState } from 'react'
import ColoredCircle from './Components/ColoredCircle.js'
import ColorText from './Components/ColorText.js'




const App = () => {

    const [color,setColor] = useState('black')

    return (
        <div>
        <ColoredCircle color={color}/>
        <h3>Type a Color!</h3> 
        <ColorText value={color} filter={(event) => setColor(event.target.value) }/>
        </div>
    )
}


export default App