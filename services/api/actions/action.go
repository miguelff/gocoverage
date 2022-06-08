package actions

import "coveragetest/pkg/transformer"

func transform(s string) string {
    return transformer.Transform(s)
}
