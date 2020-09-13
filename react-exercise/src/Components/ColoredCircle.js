import React from 'react'


const ColoredCircle = ({color}) => {
    const circleStyle = {
        height: '50px',
        width: '50px',
        backgroundColor: color,
        borderRadius: '50%',
        display: 'inline-block',
      }
    
    return (
        <div>
            <span style={circleStyle}></span>
        </div>
    )
}

export default ColoredCircle