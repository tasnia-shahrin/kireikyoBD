import React, { useState } from "react";
import "./AddProduct.css";

const AddProduct = () => {
  const [productDetails, setProductDetails] = useState({
    name: "",
    image: "",
    category: "k-beauty",
    new_price: "",
    old_price: "",
  });

  const changeHandler = (e) => {
    setProductDetails({ ...productDetails, [e.target.name]: e.target.value });
  };

  const Add_Product = async () => {
    const product = {
      name: productDetails.name.trim(),
      category: productDetails.category,
      image: productDetails.image.trim(),
      new_price: Number(productDetails.new_price),
      old_price: Number(productDetails.old_price),
    };

    if (!product.name || !product.image) {
      alert("Name and Image URL are required");
      return;
    }
    if (Number.isNaN(product.new_price) || Number.isNaN(product.old_price)) {
      alert("Prices must be numbers");
      return;
    }

    const res = await fetch("http://localhost:8080/create-products", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(product),
    });

    if (!res.ok) {
      alert("Failed to add product");
      return;
    }

    alert("Product Added ✅");

    setProductDetails({
      name: "",
      image: "",
      category: "k-beauty",
      new_price: "",
      old_price: "",
    });
  };

  return (
    <div className="add-product">
      <div className="addproduct-itemfield">
        <p>Product title</p>
        <input value={productDetails.name} onChange={changeHandler} type="text" name="name" placeholder="Type here" />
      </div>

      <div className="addproduct-itemfield">
        <p>Image URL</p>
        <input
          value={productDetails.image}
          onChange={changeHandler}
          type="text"
          name="image"
          placeholder="https://example.com/image.jpg"
        />
      </div>

      <div className="addproduct-price">
        <div className="addproduct-itemfield">
          <p>Price</p>
          <input value={productDetails.old_price} onChange={changeHandler} type="text" name="old_price" placeholder="type here" />
        </div>
        <div className="addproduct-itemfield">
          <p>Offer Price</p>
          <input value={productDetails.new_price} onChange={changeHandler} type="text" name="new_price" placeholder="type here" />
        </div>
      </div>

      <div className="addproduct-itemfield">
        <p>Product Category</p>
        <select value={productDetails.category} onChange={changeHandler} name="category" className="add-product-selector">
          <option value="k-beauty">K-Beauty</option>
          <option value="j-beauty">J-Beauty</option>
          <option value="women">Women</option>
          <option value="men">Men</option>
          <option value="kid">Kid</option>
        </select>
      </div>

      <button onClick={Add_Product} className="addproduct-btn">
        ADD
      </button>
    </div>
  );
};

export default AddProduct;
