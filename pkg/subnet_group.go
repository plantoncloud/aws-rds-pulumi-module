package pkg

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func subnetGroup(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider) (*rds.SubnetGroup, error) {
	subnetGroup, err := rds.NewSubnetGroup(ctx, "default", &rds.SubnetGroupArgs{
		Name:      pulumi.String(locals.AwsRds.Metadata.Id),
		SubnetIds: pulumi.ToStringArray(locals.AwsRds.Spec.RdsInstance.SubnetIds),
		Tags:      pulumi.ToStringMap(locals.Labels),
	}, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create subnet group")
	}

	return subnetGroup, nil
}
