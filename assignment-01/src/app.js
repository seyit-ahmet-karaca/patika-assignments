import "./styles.css";
import axios from "axios";

axios
  .get("https://my-json-server.typicode.com/modanisatech/bootcamp-db/products")
  .then((response) => {
    // Firstly, log response to the console,
    // inspect the response and see that it has data field
    console.log(response);

    // Assign data field of the response to
    // products variable below by destructuring
    // You can use alias
    const { data: products } = response;
    // Print names of all product to the console
    // by calling foreach  method (use arrow function)
    products.forEach((item) => console.log(item.name));

    // Get all products that contain "Şal" in their name (use filter method)
    // map filtered products to new object having only image and name field
    // assign mapped items to mappedProducts variable
    const mappedProducts = products
      .filter((product) => product.name.includes("Şal"))
      .map((product) => ({ name: product.name, image: product.image }));

    // Display the images and names of mappedProducts
    // You need to add them to the DOM
    // you need to use forEach method
    // You need to use flexbox
    // Position of image and text is up to you
    // You can use any style you wish
    const app = document.getElementById("app");
    mappedProducts.forEach((product) => {
      const imgTag = createImage(product.image);
      const spanTag = createName(product.name);
      const div = createDiv();
      div.appendChild(imgTag);
      div.appendChild(spanTag);
      app.appendChild(div);
    });
  });

const createImage = function (imageSource) {
  const imgTag = document.createElement("img");
  imgTag.setAttribute("src", imageSource);
  imgTag.setAttribute("class", "image");
  return imgTag;
};

const createName = function (name) {
  const spanTag = document.createElement("span");
  spanTag.textContent = name;
  spanTag.setAttribute("class", "name");
  return spanTag;
};

const createDiv = function () {
  const divTag = document.createElement("div");
  divTag.setAttribute("class", "product");
  return divTag;
};
