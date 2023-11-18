import React from 'react';
import { Route, Routes, BrowserRouter as Router } from 'react-router-dom';
import Home from './pages/home';
import Root from './pages/root';
import PaymentPage from './pages/PaymentPage';
import SubmitProblemPage from './pages/SubmitProblemPage';
import OpenRentalsPage from './pages/OpenRentalsPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Root />} />
        <Route path="/home" element={<Home />} />
        <Route path="/payment" element={<PaymentPage />} />
        <Route path="/submit-problem" element={<SubmitProblemPage />} />
        <Route path="/open-rentals" element={<OpenRentalsPage />} />
        {/* Add more routes for other pages if needed */}
      </Routes>
    </Router>
  );
}

export default App;
