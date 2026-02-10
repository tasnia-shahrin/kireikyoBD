import React, { useContext } from 'react'
import './ProductDisplay.css'
import star_icon from '../Assets/star_icon.png'
import star_dull_icon from '../Assets/star_dull_icon.png'
import { ShopContext } from '../../Context/ShopContext'

const ProductDisplay = (props) => {
    const {product}=props;
    const {addToCart}=useContext(ShopContext);
  return (
    <div className='productdisplay'>
        <div className="productdisplay-left">
            <div className="productdisplay-img-list">
                <img src={product.image} alt="" />
                <img src={product.image} alt="" />
                <img src={product.image} alt="" />
                <img src={product.image} alt="" />
            </div>
            <div className="productdisplay-img">
                <img className='productdisplay-main-img' src={product.image} alt="" />
            </div>
        </div>
        <div className="productdisplay-right">
            <h1>{product.name}</h1>
            <div className="productdisplay-right-star">
                <img src={star_icon} alt="" />
                <img src={star_icon} alt="" />
                <img src={star_icon} alt="" />
                <img src={star_icon} alt="" />
                <img src={star_dull_icon} alt="" />
                <p>(122)</p>
            </div>
            <div className="productdisplay-right-prices">
                <div className="productdisplay-right-price-old">
                    ${product.old_price}
                </div>
                <div className="productdisplay-right-price-new">
                    ${product.new_price}
                </div>
                <div className="productdisplay-right-description">
                COSRX Advanced Snail 96 Mucin Power Essence is a facial essence from the South Korean skincare brand COSRX. This essence is formulated with 96% snail mucin, which is known for its ability to hydrate, repair, and protect the skin. It also contains other beneficial ingredients such as niacinamide, which helps to brighten the skin and improve its overall tone and texture, and allantoin, which soothes and moisturizes the skin. This essence is designed to provide intense hydration and nourishment to the skin while also helping to improve its overall texture and appearance. It comes in a 100ml bottle, which is a larger size than the typical essence, making it a more cost-effective option for long-term use. It is recommended to apply a small amount to the face after cleansing and toning, before applying moisturizer. It is suitable for all skin types, especially for those with dry and dehydrated skin.

                </div>
                <div className="productdisplay-right-size">
                    <h1>Select Size</h1>
                    <div className="productdisplay-right-sizes">
                        <div>30ml</div>
                        <div>50ml</div>
                        <div>100ml</div>
                        
                    </div>
                </div>
                <button onClick={()=>{addToCart(product.id)}}>ADD TO CART</button>
                <p className="productdisplay-right-category"><span>Category : </span>Japanese Beauty,Essence</p>
                
            </div>
        </div>
    </div>
  )
}

export default ProductDisplay