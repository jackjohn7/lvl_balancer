module github.com/jackjohn7/lvl_balancer/examples

go 1.21.3

require github.com/jackjohn7/lvl_balancer/lib v0.0.0

require github.com/justinas/alice v1.2.0 // indirect

replace github.com/jackjohn7/lvl_balancer/lib v0.0.0 => ../lib
