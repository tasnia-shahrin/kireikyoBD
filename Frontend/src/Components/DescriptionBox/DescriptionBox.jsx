import React from 'react'
import './DescriptionBox.css'

const DescriptionBox = () => {
  return (
    <div className='descriptionbox'>
        <div className="descriptionbox-navigator">
            <div className="descriptionbox-nav-box">Description</div>
            <div className="descriptionbox-nav-box fade">Reviews (122)</div>
        </div>
        <div className="descriptionbox-description">
            <p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Quae quibusdam ad laudantium, rerum numquam modi provident ut deleniti asperiores ipsum, architecto deserunt impedit doloribus. Aut eligendi ut exercitationem! Exercitationem, unde!</p>
            <p>
                Lorem ipsum dolor sit amet consectetur, adipisicing elit. Corrupti ipsum, sunt nostrum eius ratione adipisci necessitatibus ad voluptatum at voluptates, nobis rem repellat earum laborum ab maxime sint iste nesciunt.
            </p>
        </div>
    </div>
  )
}

export default DescriptionBox