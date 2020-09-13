import React from 'react'

const ColorText = ({value,filter}) => {
    return (
      <input value={value} onChange={filter} />
    )
}

export default ColorText 