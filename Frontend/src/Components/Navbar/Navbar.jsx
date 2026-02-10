import React, { useContext, useRef, useState } from 'react'
import './Navbar.css'
import dropdown_icon from '../Assets/dropdown_icon.png'

import logo from '../Assets/logo.png'
import cart_icon from '../Assets/cart_icon.png'
import { Link } from 'react-router-dom'
import { ShopContext } from '../../Context/ShopContext'

export const Navbar = () => {
    const [menu,setMenu]=useState("home");
    const {getTotalCartItems}=useContext(ShopContext);
    const menuRef=useRef();

    const dropdown_toggle=(e)=>{
        menuRef.current.classList.toggle('nav-menu-visible');
        e.target.classList.toggle('open');
    }
  return (
    <div className='navbar'>
        <div className="nav-logo">
            <img src={logo} alt="" />
            <p>Kireikyo</p>
        </div>
        <img className='nav-dropdown' onClick={dropdown_toggle} src={dropdown_icon} alt="" />
        <ul ref={menuRef} className="nav-menu">
            <li onClick={()=>{setMenu("home")}}><Link style={{textDecoration: 'none'}} to='/'>Home</Link>{menu==="home"?<hr/>:<></>}</li>
            <li onClick={()=>{setMenu("j-beauty")}}><Link style={{textDecoration: 'none'}} to='/j-beauty'>J-Beauty</Link>{menu==="j-beauty"?<hr/>:<></>}</li>
            <li onClick={()=>{setMenu("k-beauty")}}><Link style={{textDecoration: 'none'}} to='/k-beauty'>K-Beauty</Link>{menu==="k-beauty"?<hr/>:<></>}</li>
            <li onClick={()=>{setMenu("hot deals")}}><Link style={{textDecoration: 'none'}} to='/hot deals'>Hot Deals</Link>{menu==="hot deals"?<hr/>:<></>}</li>
        </ul>
        <div className="nav-login-cart">
            <Link to='/login'><button>Login</button></Link>
            <Link to='/cart'><img src={cart_icon} alt="" /></Link>
            
            <div className="nav-cart-count">{getTotalCartItems()}</div>
        </div>
    </div>
  )
}
export default Navbar
