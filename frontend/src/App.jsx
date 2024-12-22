import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Layout from './components/Layout';
import LoginPage from './pages/User/LoginPage';
import RegisterPage from './pages/User/RegisterPage';
import UserInfoPage from './pages/User/UserInfoPage';
import UserAddressPage from './pages/User/UserAddressPage';
import UserLoginHistoryPage from './pages/User/UserLoginHistoryPage';
import CartAddPage from './pages/Cart/CartAddPage';
import CartListPage from './pages/Cart/CartListPage';
import OrderSubmitPage from './pages/Order/OrderSubmitPage';
import OrderDetailsPage from './pages/Order/OrderDetailsPage';
import PaymentMakePage from './pages/Payment/PaymentMakePage';
import PaymentRefundPage from './pages/Payment/PaymentRefundPage';
import PaymentStatusPage from './pages/Payment/PaymentStatusPage';
import CommentAddPage from './pages/Comment/CommentAddPage';
import CommentEditPage from './pages/Comment/CommentEditPage';
import CommentListPage from './pages/Comment/CommentListPage';
import ReplyAddPage from './pages/Reply/ReplyAddPage';
import ReplyEditPage from './pages/Reply/ReplyEditPage';
import ReplyListPage from './pages/Reply/ReplyListPage';
import SeckillEventsPage from './pages/Seckill/SeckillEventsPage';
import SeckillOrderPage from './pages/Seckill/SeckillOrderPage';
import SeckillProductsPage from './pages/Seckill/SeckillProductsPage';
import ProductAddPage from './pages/Admin/Product/ProductAddPage';
import ProductEditPage from './pages/Admin/Product/ProductEditPage';
import ProductDeletePage from './pages/Admin/Product/ProductDeletePage';
import ProductListPage from './pages/Admin/Product/ProductListPage';
import PrivateRoute from './components/PrivateRoute';

function App() {
  return (
    <Routes>
      {/* Public Routes */}
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />

      {/* Protected Routes (Requires Login) */}
      <Route element={<PrivateRoute />}>
        {/* User Pages */}
        <Route path="/user/info" element={<UserInfoPage />} />
        <Route path="/user/addresses" element={<UserAddressPage />} />
        <Route path="/user/login-history" element={<UserLoginHistoryPage />} />

        {/* Cart Pages */}
        <Route path="/cart/add" element={<CartAddPage />} />
        <Route path="/cart/list" element={<CartListPage />} />

        {/* Order Pages */}
        <Route path="/order/submit" element={<OrderSubmitPage />} />
        <Route path="/order/details/:orderId" element={<OrderDetailsPage />} />

        {/* Payment Pages */}
        <Route path="/payment/make" element={<PaymentMakePage />} />
        <Route path="/payment/refund" element={<PaymentRefundPage />} />
        <Route path="/payment/status/:paymentId" element={<PaymentStatusPage />} />

        {/* Comment Pages */}
        <Route path="/comment/add" element={<CommentAddPage />} />
        <Route path="/comment/edit/:commentId" element={<CommentEditPage />} />
        <Route path="/comment/list/:productId" element={<CommentListPage />} />

        {/* Reply Pages */}
        <Route path="/reply/add" element={<ReplyAddPage />} />
        <Route path="/reply/edit/:replyId" element={<ReplyEditPage />} />
        <Route path="/reply/list/:commentId" element={<ReplyListPage />} />

        {/* Seckill Pages */}
        <Route path="/seckill/events" element={<SeckillEventsPage />} />
        <Route path="/seckill/products" element={<SeckillProductsPage />} />
        <Route path="/seckill/order" element={<SeckillOrderPage />} />

        {/* Admin Pages (Requires Admin Role) */}
        <Route element={<PrivateRoute adminOnly={true} />}>
          <Route path="/admin/product/add" element={<ProductAddPage />} />
          <Route path="/admin/product/edit/:productId" element={<ProductEditPage />} />
          <Route path="/admin/product/delete/:productId" element={<ProductDeletePage />} />
          <Route path="/admin/product/list" element={<ProductListPage />} />
        </Route>

        {/* Home Page */}
        <Route path="/" element={<Layout><h1>Home</h1></Layout>} />
      </Route>

      {/* Catch-all route for unmatched paths */}
      <Route path="*" element={<div>404 Not Found</div>} />
    </Routes>
  );
}

export default App;