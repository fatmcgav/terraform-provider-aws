// Code generated by "internal/generate/listpages/main.go -ListOps=ListOriginAccessControls -InputPaginator=Marker -OutputPaginator=OriginAccessControlList.NextMarker -- list_origin_access_controls_pages_gen.go"; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func listOriginAccessControlsPages(ctx context.Context, conn *cloudfront.Client, input *cloudfront.ListOriginAccessControlsInput, fn func(*cloudfront.ListOriginAccessControlsOutput, bool) bool) error {
	for {
		output, err := conn.ListOriginAccessControls(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.OriginAccessControlList.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.Marker = output.OriginAccessControlList.NextMarker
	}
	return nil
}