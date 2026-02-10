
import './App.css';
import  Navbar  from './Components/Navbar/Navbar';
import  {BrowserRouter,Routes,Route}  from 'react-router-dom';
import  Home  from './Pages/Home';
import ShopCategory from './Pages/ShopCategory'
import Product from './Pages/Product'
import  LoginSignup  from './Pages/LoginSignup';
import Cart from './Pages/Cart'
import Footer from './Components/Footer/Footer';
import jbeauty_banner from './Components/Assets/banner1.jpg'
import kbeauty_banner from './Components/Assets/banner3.jpg'
import hotdeals from './Components/Assets/banner1.jpg'

function App() {
  return (
    <div>
      <BrowserRouter>
        <Navbar/>
        <Routes>
          <Route path='/' element={<Home/>}/>
          <Route path='/j-beauty' element={<ShopCategory banner={jbeauty_banner} category="j-beauty"/>}/>
          <Route path='/k-beauty' element={<ShopCategory banner={kbeauty_banner} category="k-beauty"/>}/>
          <Route path='/hot deals' element={<ShopCategory banner={hotdeals} category="hot deals"/>}/>
          <Route path='/product' element={<Product/>}>
            <Route path=':productId' element={<Product/>}/>
          </Route>
          <Route path='/cart' element={<Cart/>}/>
          <Route path='/login' element={<LoginSignup/>}/>
        </Routes>
        <Footer/>
      </BrowserRouter>
      
    </div>
  );
}

export default App;
