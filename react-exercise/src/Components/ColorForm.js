import React from 'react'

const ColorForm = ({setColor}) => {
  const changeColor = (event) => {
    event.preventDefault()
    const color  = event.target.color.value
    setColor(color)
  }
  
  return (
      <form onSubmit={changeColor}>
        <input name="color"/>
        <button type="submit">Change color!</button>
      </form>    )
}

export default ColorForm