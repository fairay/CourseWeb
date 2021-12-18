import React from "react";
import {WithParams, PP} from "../../Warpers"

class Recipe extends React.Component<PP> {
    constructor(props) {
      super(props);
    }
  
    render() {
      return(
        <div>
          <h2>Recipe {this.props.match.id}</h2>
        </div>
      )
    }  
}
const RecipeParams = () => WithParams(Recipe)
export default RecipeParams
  