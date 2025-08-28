package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) FindAll(ctx context.Context, limit int64, page int64) ([]*domain.Product, error) {
	skip := (page - 1) * limit

	pipeline := mongo.Pipeline{
		// join with var_options
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "var_options"},
			{Key: "localField", Value: "product_items.var_option_ids"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "all_var_options"},
		}}},
		// join with variations
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "variations"},
			{Key: "localField", Value: "all_var_options.variation_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "all_variations"},
		}}},
		// join with categories
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "categories"},
			{Key: "localField", Value: "category_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "category"},
		}}},
		// extract the name field from category as category_name
		{{Key: "$addFields", Value: bson.D{
			{Key: "category_name", Value: bson.D{
				{Key: "$cond", Value: bson.A{
					bson.D{{Key: "$gt", Value: bson.A{bson.D{{Key: "$size", Value: "$category"}}, 0}}},
					bson.D{{Key: "$arrayElemAt", Value: bson.A{"$category.name", 0}}},
					"", // if there is no category
				}},
			}},
		}}},

		// rebuilding product_items with their options
		{{Key: "$addFields", Value: bson.D{
			{Key: "product_items", Value: bson.D{
				{Key: "$map", Value: bson.D{
					{Key: "input", Value: "$product_items"},
					{Key: "as", Value: "item"},
					{Key: "in", Value: bson.D{
						{Key: "$mergeObjects", Value: bson.A{
							"$$item",
							bson.D{
								{Key: "options", Value: bson.D{
									{Key: "$map", Value: bson.D{
										{Key: "input", Value: bson.D{
											{Key: "$filter", Value: bson.D{
												{Key: "input", Value: "$all_var_options"},
												{Key: "as", Value: "opt"},
												{Key: "cond", Value: bson.D{
													{Key: "$in", Value: bson.A{"$$opt._id", "$$item.var_option_ids"}},
												}},
											}},
										}},
										{Key: "as", Value: "opt"},
										{Key: "in", Value: bson.D{
											{Key: "name", Value: bson.D{
												{Key: "$first", Value: bson.D{
													{Key: "$map", Value: bson.D{
														{Key: "input", Value: bson.D{
															{Key: "$filter", Value: bson.D{
																{Key: "input", Value: "$all_variations"},
																{Key: "as", Value: "v"},
																{Key: "cond", Value: bson.D{
																	{Key: "$eq", Value: bson.A{"$$v._id", "$$opt.variation_id"}},
																}},
															}},
														}},
														{Key: "as", Value: "v"},
														{Key: "in", Value: "$$v.name"},
													}},
												}},
											}},
											{Key: "var_opt_name", Value: "$$opt.label"},
											{Key: "var_opt_value", Value: "$$opt.value"},
										}},
									}},
								}},
							},
						}},
					}},
				}},
			}},
		}}},

		// pagination
		{{Key: "$skip", Value: skip}},
		{{Key: "$limit", Value: limit}},
	}

	var products []*domain.Product

	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregate pipeline: %w", err)
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("failed to decode products from cursor: %w", err)
	}

	return products, nil
}
