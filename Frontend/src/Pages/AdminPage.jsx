import React, { useState } from "react";
import './CSS/AdminPage.css'

const AdminPage = () => {
  const [product, setProduct] = useState({
    name: "",
    category: "k-beauty", // default
    image: "",
    new_price: "",
    old_price: "",
  });

  const handleChange = (e) => {
    setProduct((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    // ✅ This product.category will be exactly what you selected
    console.log("Product to send:", product);

    alert(`Added to category: ${product.category}`);

    setProduct({
      name: "",
      category: "k-beauty",
      image: "",
      new_price: "",
      old_price: "",
    });
  };

  return (
  <div className="admin-wrapper">
    <div className="admin-card">
      <h2>Admin: Add Product</h2>

      <form onSubmit={handleSubmit} className="admin-form">
        <input
          type="text"
          name="name"
          placeholder="Product Name"
          value={product.name}
          onChange={handleChange}
          required
        />

        <select
          name="category"
          value={product.category}
          onChange={handleChange}
          required
        >
          <option value="k-beauty">K-Beauty</option>
          <option value="j-beauty">J-Beauty</option>
          
        </select>

        <input
          type="text"
          name="image"
          placeholder="Image URL"
          value={product.image}
          onChange={handleChange}
          required
        />

        <input
          type="number"
          name="new_price"
          placeholder="New Price"
          value={product.new_price}
          onChange={handleChange}
          required
        />

        <input
          type="number"
          name="old_price"
          placeholder="Old Price"
          value={product.old_price}
          onChange={handleChange}
          required
        />

        <button type="submit" className="admin-btn">
          Add Product
        </button>
      </form>
    </div>
  </div>
);

};



export default AdminPage;
