import React, {useEffect, useState} from "react";
import {BrowserRouter as Router} from "react-router-dom";
import "./Posts.css";
import Post from "../Post/Post";
import axios from "axios"


function Posts() {
	const [blogs, setBlogs] = useState(null);
	const [error, setErrors] = useState(null);
	const [loading, setLoading] = useState(true);
	const fetchBlogs = async ()=> {
	try {
		setLoading(true)
			const res = await axios.get("http://localhost:8080/api/blogs");
			setBlogs(res.data);
			setErrors(null);
		} catch(e){
		setErrors(e);
		setBlogs(null);
	}
	setLoading(false);
	
	};
	useEffect(() => {
		fetchBlogs();
	},[]);
	if(loading){
		return <div className="loader">Loading...</div>
	} else if (error){
		console.log(error)
		return <div className="error">oops.. please reload the page</div>
	}else if(blogs.length === 0){
		return <div>No Blogs available</div>
	}else{
		return ( <div className="blogs-container">
	{blogs.map((blog, index) => (
		<Post key={index} index={index} post={blog} />
	))}
	</div> 
	); 
	}
	
}

export default Posts;
