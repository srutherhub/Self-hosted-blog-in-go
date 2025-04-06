package constants

const HtmlStyle string = `<style>
body {
background-color:#181C14;
color:#ECDFCC;
font-family:sans-serif;
display: flex;
flex-direction:column;
align-items:center;
} 

article {
line-height:1.6;
max-width:36rem;
}

.links-container {
	display:flex;
	flex-direction:column;
}

img {
max-width:36rem;
}
h1,h2,h3 {
padding:0.5rem;
margin:0
}
a {
color:#ECDFCC;
font-family:sans-serif;
text-decoration:none;
}

@media (max-width: 600px) { 
  article {
    max-width: 95%; 
  }
		img {
max-width:95%;
}
}

</style>`
