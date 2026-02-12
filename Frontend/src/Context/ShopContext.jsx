import { createContext, useEffect, useMemo, useState } from "react";

export const ShopContext = createContext(null);

const ShopContextProvider = (props) => {
  const [all_product, setAllProduct] = useState([]);
  const [cartItems, setCartItems] = useState({});

  const fetchProducts = async () => {
    const res = await fetch("http://localhost:8080/products");
    const data = await res.json();
    setAllProduct(data);

    // initialize cart keys by product IDs (only first time)
    setCartItems((prev) => {
      if (Object.keys(prev).length > 0) return prev;
      const cart = {};
      for (const p of data) cart[p.id] = 0;
      return cart;
    });
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  const addToCart = (itemId) => {
    setCartItems((prev) => ({ ...prev, [itemId]: (prev[itemId] || 0) + 1 }));
  };

  const removeFromCart = (itemId) => {
    setCartItems((prev) => ({
      ...prev,
      [itemId]: Math.max((prev[itemId] || 0) - 1, 0),
    }));
  };

  const getTotalCartAmount = () => {
    let totalAmount = 0;
    for (const id in cartItems) {
      const qty = cartItems[id];
      if (qty > 0) {
        const itemInfo = all_product.find((p) => p.id === Number(id));
        if (itemInfo) totalAmount += itemInfo.new_price * qty;
      }
    }
    return totalAmount;
  };

  const getTotalCartItems = () => {
    let total = 0;
    for (const id in cartItems) total += cartItems[id];
    return total;
  };

  const contextValue = useMemo(
    () => ({
      all_product,
      fetchProducts,
      cartItems,
      addToCart,
      removeFromCart,
      getTotalCartAmount,
      getTotalCartItems,
    }),
    [all_product, cartItems]
  );

  return (
    <ShopContext.Provider value={contextValue}>
      {props.children}
    </ShopContext.Provider>
  );
};

export default ShopContextProvider;
