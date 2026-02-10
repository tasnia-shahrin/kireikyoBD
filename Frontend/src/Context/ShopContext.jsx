import { createContext, useEffect, useMemo, useState } from "react";

export const ShopContext = createContext(null);

const ShopContextProvider = (props) => {
  const [all_product, setAllProduct] = useState([]);
  const [cartItems, setCartItems] = useState({});

  useEffect(() => {
    fetch("http://localhost:8080/products")
      .then((res) => res.json())
      .then((data) => {
        setAllProduct(data);

        const cart = {};
        for (const p of data) cart[p.id] = 0;
        setCartItems(cart);
      })
      .catch((err) => console.error("Failed to load products:", err));
  }, []);

  const addToCart = (itemId) => {
    setCartItems((prev) => ({
      ...prev,
      [itemId]: (prev[itemId] || 0) + 1,
    }));
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
    let totalItem = 0;
    for (const id in cartItems) totalItem += cartItems[id];
    return totalItem;
  };

  const contextValue = useMemo(
    () => ({
      getTotalCartItems,
      getTotalCartAmount,
      all_product,
      cartItems,
      addToCart,
      removeFromCart,
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
