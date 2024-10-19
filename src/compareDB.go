package main

import (
	"fmt"
	"sort"
)


func SortByNameAndIngredients(data *BakeryData) {
	sort.SliceStable(data.Cakes, func(i, j int) bool {
		if data.Cakes[i].Name != data.Cakes[j].Name {
			return data.Cakes[i].Name < data.Cakes[j].Name
		}
		return len(data.Cakes[i].Ingredients) < len(data.Cakes[j].Ingredients)
	})
}


func compareDB(oldData BakeryData,newData BakeryData) {
		SortByNameAndIngredients(&newData)
		SortByNameAndIngredients(&oldData)
		for i:=0;i<len(oldData.Cakes);i++{
			
			if newData.Cakes[i].Name==oldData.Cakes[i].Name{
				if newData.Cakes[i].Time!=oldData.Cakes[i].Time{
					fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",oldData.Cakes[i].Name,newData.Cakes[i].Time,oldData.Cakes[i].Time)

				}
				for j:=0;j<len(oldData.Cakes[i].Ingredients);j++{
					inflag:=false
					for k:=0;k<len(newData.Cakes[i].Ingredients);k++{
						if oldData.Cakes[i].Ingredients[j].IngredientName==newData.Cakes[i].Ingredients[k].IngredientName{
							inflag=true
							
							if oldData.Cakes[i].Ingredients[j].IngredientCount!=newData.Cakes[i].Ingredients[k].IngredientCount{
						
								fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",oldData.Cakes[i].Ingredients[j].IngredientName,newData.Cakes[i].Name,oldData.Cakes[i].Ingredients[j].IngredientCount,newData.Cakes[i].Ingredients[k].IngredientCount)
							}
							if oldData.Cakes[i].Ingredients[j].IngredientUnit!=newData.Cakes[i].Ingredients[k].IngredientUnit {
								if newData.Cakes[i].Ingredients[k].IngredientUnit==""{
									fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", oldData.Cakes[i].Ingredients[k].IngredientUnit,oldData.Cakes[i].Ingredients[j].IngredientName,newData.Cakes[i].Name)
								} else if  oldData.Cakes[i].Ingredients[k].IngredientUnit==""{
									fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", newData.Cakes[i].Ingredients[k].IngredientUnit,oldData.Cakes[i].Ingredients[j].IngredientName,newData.Cakes[i].Name)

								}else{
									fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",oldData.Cakes[i].Ingredients[j].IngredientName,newData.Cakes[i].Name,oldData.Cakes[i].Ingredients[j].IngredientUnit,newData.Cakes[i].Ingredients[k].IngredientUnit)

								}
							}
						}
					}
					if !inflag{
						fmt.Printf("REMOVED ingredient \"%s\" for cake  \"%s\"\n",oldData.Cakes[i].Ingredients[j].IngredientName, newData.Cakes[i].Name)
					}
				}
				for j:=0;j<len(newData.Cakes[i].Ingredients);j++{
					inflag:=false
					for k:=0;k<len(oldData.Cakes[i].Ingredients);k++{
						if oldData.Cakes[i].Ingredients[k].IngredientName==newData.Cakes[i].Ingredients[j].IngredientName{
							inflag=true
						}
					}
					if !inflag{
						fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\"\n", newData.Cakes[i].Ingredients[j].IngredientName,newData.Cakes[i].Name )
					}
				}
			}else{
				fmt.Printf("ADDED cake \"%s\"\n", newData.Cakes[i].Name)
				fmt.Printf("REMOVED cake \"%s\"\n", oldData.Cakes[i].Name)
			}
			
		}
	
	
}