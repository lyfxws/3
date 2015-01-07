package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"sync"
	"unsafe"
)

// CUDA handle for addcubicanisotropy kernel
var addcubicanisotropy_code cu.Function

// Stores the arguments for addcubicanisotropy kernel invocation
type addcubicanisotropy_args_t struct {
	arg_Bx      unsafe.Pointer
	arg_By      unsafe.Pointer
	arg_Bz      unsafe.Pointer
	arg_mx      unsafe.Pointer
	arg_my      unsafe.Pointer
	arg_mz      unsafe.Pointer
	arg_K1LUT   unsafe.Pointer
	arg_K2LUT   unsafe.Pointer
	arg_K3LUT   unsafe.Pointer
	arg_C1xLUT  unsafe.Pointer
	arg_C1yLUT  unsafe.Pointer
	arg_C1zLUT  unsafe.Pointer
	arg_C2xLUT  unsafe.Pointer
	arg_C2yLUT  unsafe.Pointer
	arg_C2zLUT  unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_N       int
	argptr      [17]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for addcubicanisotropy kernel invocation
var addcubicanisotropy_args addcubicanisotropy_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	addcubicanisotropy_args.argptr[0] = unsafe.Pointer(&addcubicanisotropy_args.arg_Bx)
	addcubicanisotropy_args.argptr[1] = unsafe.Pointer(&addcubicanisotropy_args.arg_By)
	addcubicanisotropy_args.argptr[2] = unsafe.Pointer(&addcubicanisotropy_args.arg_Bz)
	addcubicanisotropy_args.argptr[3] = unsafe.Pointer(&addcubicanisotropy_args.arg_mx)
	addcubicanisotropy_args.argptr[4] = unsafe.Pointer(&addcubicanisotropy_args.arg_my)
	addcubicanisotropy_args.argptr[5] = unsafe.Pointer(&addcubicanisotropy_args.arg_mz)
	addcubicanisotropy_args.argptr[6] = unsafe.Pointer(&addcubicanisotropy_args.arg_K1LUT)
	addcubicanisotropy_args.argptr[7] = unsafe.Pointer(&addcubicanisotropy_args.arg_K2LUT)
	addcubicanisotropy_args.argptr[8] = unsafe.Pointer(&addcubicanisotropy_args.arg_K3LUT)
	addcubicanisotropy_args.argptr[9] = unsafe.Pointer(&addcubicanisotropy_args.arg_C1xLUT)
	addcubicanisotropy_args.argptr[10] = unsafe.Pointer(&addcubicanisotropy_args.arg_C1yLUT)
	addcubicanisotropy_args.argptr[11] = unsafe.Pointer(&addcubicanisotropy_args.arg_C1zLUT)
	addcubicanisotropy_args.argptr[12] = unsafe.Pointer(&addcubicanisotropy_args.arg_C2xLUT)
	addcubicanisotropy_args.argptr[13] = unsafe.Pointer(&addcubicanisotropy_args.arg_C2yLUT)
	addcubicanisotropy_args.argptr[14] = unsafe.Pointer(&addcubicanisotropy_args.arg_C2zLUT)
	addcubicanisotropy_args.argptr[15] = unsafe.Pointer(&addcubicanisotropy_args.arg_regions)
	addcubicanisotropy_args.argptr[16] = unsafe.Pointer(&addcubicanisotropy_args.arg_N)
}

// Wrapper for addcubicanisotropy CUDA kernel, asynchronous.
func k_addcubicanisotropy_async(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, K1LUT unsafe.Pointer, K2LUT unsafe.Pointer, K3LUT unsafe.Pointer, C1xLUT unsafe.Pointer, C1yLUT unsafe.Pointer, C1zLUT unsafe.Pointer, C2xLUT unsafe.Pointer, C2yLUT unsafe.Pointer, C2zLUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	addcubicanisotropy_args.Lock()
	defer addcubicanisotropy_args.Unlock()

	if addcubicanisotropy_code == 0 {
		addcubicanisotropy_code = fatbinLoad(addcubicanisotropy_map, "addcubicanisotropy")
	}

	addcubicanisotropy_args.arg_Bx = Bx
	addcubicanisotropy_args.arg_By = By
	addcubicanisotropy_args.arg_Bz = Bz
	addcubicanisotropy_args.arg_mx = mx
	addcubicanisotropy_args.arg_my = my
	addcubicanisotropy_args.arg_mz = mz
	addcubicanisotropy_args.arg_K1LUT = K1LUT
	addcubicanisotropy_args.arg_K2LUT = K2LUT
	addcubicanisotropy_args.arg_K3LUT = K3LUT
	addcubicanisotropy_args.arg_C1xLUT = C1xLUT
	addcubicanisotropy_args.arg_C1yLUT = C1yLUT
	addcubicanisotropy_args.arg_C1zLUT = C1zLUT
	addcubicanisotropy_args.arg_C2xLUT = C2xLUT
	addcubicanisotropy_args.arg_C2yLUT = C2yLUT
	addcubicanisotropy_args.arg_C2zLUT = C2zLUT
	addcubicanisotropy_args.arg_regions = regions
	addcubicanisotropy_args.arg_N = N

	args := addcubicanisotropy_args.argptr[:]
	cu.LaunchKernel(addcubicanisotropy_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for addcubicanisotropy kernel.
var addcubicanisotropy_map = map[int]string{0: "",
	20: addcubicanisotropy_ptx_20,
	30: addcubicanisotropy_ptx_30,
	35: addcubicanisotropy_ptx_35}

// addcubicanisotropy PTX code for various compute capabilities.
const (
	addcubicanisotropy_ptx_20 = `
.version 4.0
.target sm_20
.address_size 64


.visible .entry addcubicanisotropy(
	.param .u64 addcubicanisotropy_param_0,
	.param .u64 addcubicanisotropy_param_1,
	.param .u64 addcubicanisotropy_param_2,
	.param .u64 addcubicanisotropy_param_3,
	.param .u64 addcubicanisotropy_param_4,
	.param .u64 addcubicanisotropy_param_5,
	.param .u64 addcubicanisotropy_param_6,
	.param .u64 addcubicanisotropy_param_7,
	.param .u64 addcubicanisotropy_param_8,
	.param .u64 addcubicanisotropy_param_9,
	.param .u64 addcubicanisotropy_param_10,
	.param .u64 addcubicanisotropy_param_11,
	.param .u64 addcubicanisotropy_param_12,
	.param .u64 addcubicanisotropy_param_13,
	.param .u64 addcubicanisotropy_param_14,
	.param .u64 addcubicanisotropy_param_15,
	.param .u32 addcubicanisotropy_param_16
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<16>;
	.reg .f32 	%f<139>;
	.reg .s64 	%rd<54>;


	ld.param.u64 	%rd2, [addcubicanisotropy_param_0];
	ld.param.u64 	%rd3, [addcubicanisotropy_param_1];
	ld.param.u64 	%rd4, [addcubicanisotropy_param_2];
	ld.param.u64 	%rd5, [addcubicanisotropy_param_3];
	ld.param.u64 	%rd6, [addcubicanisotropy_param_4];
	ld.param.u64 	%rd7, [addcubicanisotropy_param_5];
	ld.param.u64 	%rd8, [addcubicanisotropy_param_6];
	ld.param.u64 	%rd9, [addcubicanisotropy_param_7];
	ld.param.u64 	%rd10, [addcubicanisotropy_param_8];
	ld.param.u64 	%rd11, [addcubicanisotropy_param_9];
	ld.param.u64 	%rd12, [addcubicanisotropy_param_10];
	ld.param.u64 	%rd13, [addcubicanisotropy_param_11];
	ld.param.u64 	%rd14, [addcubicanisotropy_param_12];
	ld.param.u64 	%rd15, [addcubicanisotropy_param_13];
	ld.param.u64 	%rd16, [addcubicanisotropy_param_14];
	ld.param.u64 	%rd17, [addcubicanisotropy_param_15];
	ld.param.u32 	%r2, [addcubicanisotropy_param_16];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_8;

	cvta.to.global.u64 	%rd18, %rd17;
	cvt.s64.s32	%rd19, %r1;
	add.s64 	%rd20, %rd18, %rd19;
	ld.global.u8 	%rd1, [%rd20];
	cvta.to.global.u64 	%rd21, %rd8;
	shl.b64 	%rd22, %rd1, 2;
	add.s64 	%rd23, %rd21, %rd22;
	ld.global.f32 	%f1, [%rd23];
	cvta.to.global.u64 	%rd24, %rd9;
	add.s64 	%rd25, %rd24, %rd22;
	ld.global.f32 	%f2, [%rd25];
	cvta.to.global.u64 	%rd26, %rd10;
	add.s64 	%rd27, %rd26, %rd22;
	ld.global.f32 	%f3, [%rd27];
	cvta.to.global.u64 	%rd28, %rd11;
	add.s64 	%rd29, %rd28, %rd22;
	cvta.to.global.u64 	%rd30, %rd12;
	add.s64 	%rd31, %rd30, %rd22;
	cvta.to.global.u64 	%rd32, %rd13;
	add.s64 	%rd33, %rd32, %rd22;
	ld.global.f32 	%f4, [%rd29];
	ld.global.f32 	%f5, [%rd31];
	mul.f32 	%f19, %f5, %f5;
	fma.rn.f32 	%f20, %f4, %f4, %f19;
	ld.global.f32 	%f6, [%rd33];
	fma.rn.f32 	%f21, %f6, %f6, %f20;
	sqrt.rn.f32 	%f7, %f21;
	setp.neu.f32	%p2, %f7, 0f00000000;
	@%p2 bra 	BB0_3;

	mov.f32 	%f137, 0f00000000;
	bra.uni 	BB0_4;

BB0_3:
	rcp.rn.f32 	%f137, %f7;

BB0_4:
	mul.f32 	%f10, %f137, %f4;
	mul.f32 	%f11, %f137, %f5;
	mul.f32 	%f12, %f137, %f6;
	cvta.to.global.u64 	%rd34, %rd14;
	add.s64 	%rd36, %rd34, %rd22;
	cvta.to.global.u64 	%rd37, %rd15;
	add.s64 	%rd38, %rd37, %rd22;
	cvta.to.global.u64 	%rd39, %rd16;
	add.s64 	%rd40, %rd39, %rd22;
	ld.global.f32 	%f13, [%rd36];
	ld.global.f32 	%f14, [%rd38];
	mul.f32 	%f23, %f14, %f14;
	fma.rn.f32 	%f24, %f13, %f13, %f23;
	ld.global.f32 	%f15, [%rd40];
	fma.rn.f32 	%f25, %f15, %f15, %f24;
	sqrt.rn.f32 	%f16, %f25;
	setp.neu.f32	%p3, %f16, 0f00000000;
	@%p3 bra 	BB0_6;

	mov.f32 	%f138, 0f00000000;
	bra.uni 	BB0_7;

BB0_6:
	rcp.rn.f32 	%f138, %f16;

BB0_7:
	mul.f32 	%f27, %f138, %f15;
	mul.f32 	%f28, %f11, %f27;
	mul.f32 	%f29, %f138, %f14;
	mul.f32 	%f30, %f12, %f29;
	sub.f32 	%f31, %f28, %f30;
	mul.f32 	%f32, %f138, %f13;
	mul.f32 	%f33, %f12, %f32;
	mul.f32 	%f34, %f10, %f27;
	sub.f32 	%f35, %f33, %f34;
	mul.f32 	%f36, %f10, %f29;
	mul.f32 	%f37, %f11, %f32;
	sub.f32 	%f38, %f36, %f37;
	cvta.to.global.u64 	%rd41, %rd5;
	mul.wide.s32 	%rd42, %r1, 4;
	add.s64 	%rd43, %rd41, %rd42;
	cvta.to.global.u64 	%rd44, %rd6;
	add.s64 	%rd45, %rd44, %rd42;
	cvta.to.global.u64 	%rd46, %rd7;
	add.s64 	%rd47, %rd46, %rd42;
	ld.global.f32 	%f39, [%rd43];
	ld.global.f32 	%f40, [%rd45];
	mul.f32 	%f41, %f11, %f40;
	fma.rn.f32 	%f42, %f10, %f39, %f41;
	ld.global.f32 	%f43, [%rd47];
	fma.rn.f32 	%f44, %f12, %f43, %f42;
	mul.f32 	%f45, %f29, %f40;
	fma.rn.f32 	%f46, %f32, %f39, %f45;
	fma.rn.f32 	%f47, %f27, %f43, %f46;
	mul.f32 	%f48, %f35, %f40;
	fma.rn.f32 	%f49, %f31, %f39, %f48;
	fma.rn.f32 	%f50, %f38, %f43, %f49;
	mul.f32 	%f51, %f47, %f47;
	mul.f32 	%f52, %f50, %f50;
	add.f32 	%f53, %f51, %f52;
	mul.f32 	%f54, %f44, %f10;
	mul.f32 	%f55, %f44, %f11;
	mul.f32 	%f56, %f44, %f12;
	mul.f32 	%f57, %f44, %f44;
	add.f32 	%f58, %f57, %f52;
	mul.f32 	%f59, %f47, %f32;
	mul.f32 	%f60, %f47, %f29;
	mul.f32 	%f61, %f47, %f27;
	mul.f32 	%f62, %f58, %f59;
	mul.f32 	%f63, %f58, %f60;
	mul.f32 	%f64, %f58, %f61;
	fma.rn.f32 	%f65, %f53, %f54, %f62;
	fma.rn.f32 	%f66, %f53, %f55, %f63;
	fma.rn.f32 	%f67, %f53, %f56, %f64;
	add.f32 	%f68, %f57, %f51;
	mul.f32 	%f69, %f50, %f31;
	mul.f32 	%f70, %f50, %f35;
	mul.f32 	%f71, %f50, %f38;
	fma.rn.f32 	%f72, %f68, %f69, %f65;
	fma.rn.f32 	%f73, %f68, %f70, %f66;
	fma.rn.f32 	%f74, %f68, %f71, %f67;
	mul.f32 	%f75, %f1, 0fC0000000;
	mul.f32 	%f76, %f75, %f72;
	mul.f32 	%f77, %f75, %f73;
	mul.f32 	%f78, %f75, %f74;
	mul.f32 	%f79, %f51, %f52;
	mul.f32 	%f80, %f57, %f52;
	mul.f32 	%f81, %f80, %f59;
	mul.f32 	%f82, %f80, %f60;
	mul.f32 	%f83, %f80, %f61;
	fma.rn.f32 	%f84, %f79, %f54, %f81;
	fma.rn.f32 	%f85, %f79, %f55, %f82;
	fma.rn.f32 	%f86, %f79, %f56, %f83;
	mul.f32 	%f87, %f57, %f51;
	fma.rn.f32 	%f88, %f87, %f69, %f84;
	fma.rn.f32 	%f89, %f87, %f70, %f85;
	fma.rn.f32 	%f90, %f87, %f71, %f86;
	add.f32 	%f91, %f2, %f2;
	mul.f32 	%f92, %f91, %f88;
	mul.f32 	%f93, %f91, %f89;
	mul.f32 	%f94, %f91, %f90;
	sub.f32 	%f95, %f76, %f92;
	sub.f32 	%f96, %f77, %f93;
	sub.f32 	%f97, %f78, %f94;
	mul.f32 	%f98, %f51, %f51;
	mul.f32 	%f99, %f52, %f52;
	fma.rn.f32 	%f100, %f52, %f52, %f98;
	mul.f32 	%f101, %f57, %f44;
	mul.f32 	%f102, %f101, %f10;
	mul.f32 	%f103, %f101, %f11;
	mul.f32 	%f104, %f101, %f12;
	fma.rn.f32 	%f105, %f57, %f57, %f99;
	mul.f32 	%f106, %f51, %f47;
	mul.f32 	%f107, %f106, %f32;
	mul.f32 	%f108, %f106, %f29;
	mul.f32 	%f109, %f106, %f27;
	mul.f32 	%f110, %f105, %f107;
	mul.f32 	%f111, %f105, %f108;
	mul.f32 	%f112, %f105, %f109;
	fma.rn.f32 	%f113, %f100, %f102, %f110;
	fma.rn.f32 	%f114, %f100, %f103, %f111;
	fma.rn.f32 	%f115, %f100, %f104, %f112;
	fma.rn.f32 	%f116, %f57, %f57, %f98;
	mul.f32 	%f117, %f52, %f50;
	mul.f32 	%f118, %f117, %f31;
	mul.f32 	%f119, %f117, %f35;
	mul.f32 	%f120, %f117, %f38;
	fma.rn.f32 	%f121, %f116, %f118, %f113;
	fma.rn.f32 	%f122, %f116, %f119, %f114;
	fma.rn.f32 	%f123, %f116, %f120, %f115;
	mul.f32 	%f124, %f3, 0f40800000;
	mul.f32 	%f125, %f124, %f121;
	mul.f32 	%f126, %f124, %f122;
	mul.f32 	%f127, %f124, %f123;
	sub.f32 	%f128, %f95, %f125;
	sub.f32 	%f129, %f96, %f126;
	sub.f32 	%f130, %f97, %f127;
	cvta.to.global.u64 	%rd48, %rd2;
	add.s64 	%rd49, %rd48, %rd42;
	ld.global.f32 	%f131, [%rd49];
	add.f32 	%f132, %f131, %f128;
	st.global.f32 	[%rd49], %f132;
	cvta.to.global.u64 	%rd50, %rd3;
	add.s64 	%rd51, %rd50, %rd42;
	ld.global.f32 	%f133, [%rd51];
	add.f32 	%f134, %f133, %f129;
	st.global.f32 	[%rd51], %f134;
	cvta.to.global.u64 	%rd52, %rd4;
	add.s64 	%rd53, %rd52, %rd42;
	ld.global.f32 	%f135, [%rd53];
	add.f32 	%f136, %f135, %f130;
	st.global.f32 	[%rd53], %f136;

BB0_8:
	ret;
}


`
	addcubicanisotropy_ptx_30 = `
.version 4.0
.target sm_30
.address_size 64


.visible .entry addcubicanisotropy(
	.param .u64 addcubicanisotropy_param_0,
	.param .u64 addcubicanisotropy_param_1,
	.param .u64 addcubicanisotropy_param_2,
	.param .u64 addcubicanisotropy_param_3,
	.param .u64 addcubicanisotropy_param_4,
	.param .u64 addcubicanisotropy_param_5,
	.param .u64 addcubicanisotropy_param_6,
	.param .u64 addcubicanisotropy_param_7,
	.param .u64 addcubicanisotropy_param_8,
	.param .u64 addcubicanisotropy_param_9,
	.param .u64 addcubicanisotropy_param_10,
	.param .u64 addcubicanisotropy_param_11,
	.param .u64 addcubicanisotropy_param_12,
	.param .u64 addcubicanisotropy_param_13,
	.param .u64 addcubicanisotropy_param_14,
	.param .u64 addcubicanisotropy_param_15,
	.param .u32 addcubicanisotropy_param_16
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<16>;
	.reg .f32 	%f<139>;
	.reg .s64 	%rd<54>;


	ld.param.u64 	%rd2, [addcubicanisotropy_param_0];
	ld.param.u64 	%rd3, [addcubicanisotropy_param_1];
	ld.param.u64 	%rd4, [addcubicanisotropy_param_2];
	ld.param.u64 	%rd5, [addcubicanisotropy_param_3];
	ld.param.u64 	%rd6, [addcubicanisotropy_param_4];
	ld.param.u64 	%rd7, [addcubicanisotropy_param_5];
	ld.param.u64 	%rd8, [addcubicanisotropy_param_6];
	ld.param.u64 	%rd9, [addcubicanisotropy_param_7];
	ld.param.u64 	%rd10, [addcubicanisotropy_param_8];
	ld.param.u64 	%rd11, [addcubicanisotropy_param_9];
	ld.param.u64 	%rd12, [addcubicanisotropy_param_10];
	ld.param.u64 	%rd13, [addcubicanisotropy_param_11];
	ld.param.u64 	%rd14, [addcubicanisotropy_param_12];
	ld.param.u64 	%rd15, [addcubicanisotropy_param_13];
	ld.param.u64 	%rd16, [addcubicanisotropy_param_14];
	ld.param.u64 	%rd17, [addcubicanisotropy_param_15];
	ld.param.u32 	%r2, [addcubicanisotropy_param_16];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_8;

	cvta.to.global.u64 	%rd18, %rd17;
	cvt.s64.s32	%rd19, %r1;
	add.s64 	%rd20, %rd18, %rd19;
	ld.global.u8 	%rd1, [%rd20];
	cvta.to.global.u64 	%rd21, %rd8;
	shl.b64 	%rd22, %rd1, 2;
	add.s64 	%rd23, %rd21, %rd22;
	ld.global.f32 	%f1, [%rd23];
	cvta.to.global.u64 	%rd24, %rd9;
	add.s64 	%rd25, %rd24, %rd22;
	ld.global.f32 	%f2, [%rd25];
	cvta.to.global.u64 	%rd26, %rd10;
	add.s64 	%rd27, %rd26, %rd22;
	ld.global.f32 	%f3, [%rd27];
	cvta.to.global.u64 	%rd28, %rd11;
	add.s64 	%rd29, %rd28, %rd22;
	cvta.to.global.u64 	%rd30, %rd12;
	add.s64 	%rd31, %rd30, %rd22;
	cvta.to.global.u64 	%rd32, %rd13;
	add.s64 	%rd33, %rd32, %rd22;
	ld.global.f32 	%f4, [%rd29];
	ld.global.f32 	%f5, [%rd31];
	mul.f32 	%f19, %f5, %f5;
	fma.rn.f32 	%f20, %f4, %f4, %f19;
	ld.global.f32 	%f6, [%rd33];
	fma.rn.f32 	%f21, %f6, %f6, %f20;
	sqrt.rn.f32 	%f7, %f21;
	setp.neu.f32	%p2, %f7, 0f00000000;
	@%p2 bra 	BB0_3;

	mov.f32 	%f137, 0f00000000;
	bra.uni 	BB0_4;

BB0_3:
	rcp.rn.f32 	%f137, %f7;

BB0_4:
	mul.f32 	%f10, %f137, %f4;
	mul.f32 	%f11, %f137, %f5;
	mul.f32 	%f12, %f137, %f6;
	cvta.to.global.u64 	%rd34, %rd14;
	add.s64 	%rd36, %rd34, %rd22;
	cvta.to.global.u64 	%rd37, %rd15;
	add.s64 	%rd38, %rd37, %rd22;
	cvta.to.global.u64 	%rd39, %rd16;
	add.s64 	%rd40, %rd39, %rd22;
	ld.global.f32 	%f13, [%rd36];
	ld.global.f32 	%f14, [%rd38];
	mul.f32 	%f23, %f14, %f14;
	fma.rn.f32 	%f24, %f13, %f13, %f23;
	ld.global.f32 	%f15, [%rd40];
	fma.rn.f32 	%f25, %f15, %f15, %f24;
	sqrt.rn.f32 	%f16, %f25;
	setp.neu.f32	%p3, %f16, 0f00000000;
	@%p3 bra 	BB0_6;

	mov.f32 	%f138, 0f00000000;
	bra.uni 	BB0_7;

BB0_6:
	rcp.rn.f32 	%f138, %f16;

BB0_7:
	mul.f32 	%f27, %f138, %f15;
	mul.f32 	%f28, %f11, %f27;
	mul.f32 	%f29, %f138, %f14;
	mul.f32 	%f30, %f12, %f29;
	sub.f32 	%f31, %f28, %f30;
	mul.f32 	%f32, %f138, %f13;
	mul.f32 	%f33, %f12, %f32;
	mul.f32 	%f34, %f10, %f27;
	sub.f32 	%f35, %f33, %f34;
	mul.f32 	%f36, %f10, %f29;
	mul.f32 	%f37, %f11, %f32;
	sub.f32 	%f38, %f36, %f37;
	cvta.to.global.u64 	%rd41, %rd5;
	mul.wide.s32 	%rd42, %r1, 4;
	add.s64 	%rd43, %rd41, %rd42;
	cvta.to.global.u64 	%rd44, %rd6;
	add.s64 	%rd45, %rd44, %rd42;
	cvta.to.global.u64 	%rd46, %rd7;
	add.s64 	%rd47, %rd46, %rd42;
	ld.global.f32 	%f39, [%rd43];
	ld.global.f32 	%f40, [%rd45];
	mul.f32 	%f41, %f11, %f40;
	fma.rn.f32 	%f42, %f10, %f39, %f41;
	ld.global.f32 	%f43, [%rd47];
	fma.rn.f32 	%f44, %f12, %f43, %f42;
	mul.f32 	%f45, %f29, %f40;
	fma.rn.f32 	%f46, %f32, %f39, %f45;
	fma.rn.f32 	%f47, %f27, %f43, %f46;
	mul.f32 	%f48, %f35, %f40;
	fma.rn.f32 	%f49, %f31, %f39, %f48;
	fma.rn.f32 	%f50, %f38, %f43, %f49;
	mul.f32 	%f51, %f47, %f47;
	mul.f32 	%f52, %f50, %f50;
	add.f32 	%f53, %f51, %f52;
	mul.f32 	%f54, %f44, %f10;
	mul.f32 	%f55, %f44, %f11;
	mul.f32 	%f56, %f44, %f12;
	mul.f32 	%f57, %f44, %f44;
	add.f32 	%f58, %f57, %f52;
	mul.f32 	%f59, %f47, %f32;
	mul.f32 	%f60, %f47, %f29;
	mul.f32 	%f61, %f47, %f27;
	mul.f32 	%f62, %f58, %f59;
	mul.f32 	%f63, %f58, %f60;
	mul.f32 	%f64, %f58, %f61;
	fma.rn.f32 	%f65, %f53, %f54, %f62;
	fma.rn.f32 	%f66, %f53, %f55, %f63;
	fma.rn.f32 	%f67, %f53, %f56, %f64;
	add.f32 	%f68, %f57, %f51;
	mul.f32 	%f69, %f50, %f31;
	mul.f32 	%f70, %f50, %f35;
	mul.f32 	%f71, %f50, %f38;
	fma.rn.f32 	%f72, %f68, %f69, %f65;
	fma.rn.f32 	%f73, %f68, %f70, %f66;
	fma.rn.f32 	%f74, %f68, %f71, %f67;
	mul.f32 	%f75, %f1, 0fC0000000;
	mul.f32 	%f76, %f75, %f72;
	mul.f32 	%f77, %f75, %f73;
	mul.f32 	%f78, %f75, %f74;
	mul.f32 	%f79, %f51, %f52;
	mul.f32 	%f80, %f57, %f52;
	mul.f32 	%f81, %f80, %f59;
	mul.f32 	%f82, %f80, %f60;
	mul.f32 	%f83, %f80, %f61;
	fma.rn.f32 	%f84, %f79, %f54, %f81;
	fma.rn.f32 	%f85, %f79, %f55, %f82;
	fma.rn.f32 	%f86, %f79, %f56, %f83;
	mul.f32 	%f87, %f57, %f51;
	fma.rn.f32 	%f88, %f87, %f69, %f84;
	fma.rn.f32 	%f89, %f87, %f70, %f85;
	fma.rn.f32 	%f90, %f87, %f71, %f86;
	add.f32 	%f91, %f2, %f2;
	mul.f32 	%f92, %f91, %f88;
	mul.f32 	%f93, %f91, %f89;
	mul.f32 	%f94, %f91, %f90;
	sub.f32 	%f95, %f76, %f92;
	sub.f32 	%f96, %f77, %f93;
	sub.f32 	%f97, %f78, %f94;
	mul.f32 	%f98, %f51, %f51;
	mul.f32 	%f99, %f52, %f52;
	fma.rn.f32 	%f100, %f52, %f52, %f98;
	mul.f32 	%f101, %f57, %f44;
	mul.f32 	%f102, %f101, %f10;
	mul.f32 	%f103, %f101, %f11;
	mul.f32 	%f104, %f101, %f12;
	fma.rn.f32 	%f105, %f57, %f57, %f99;
	mul.f32 	%f106, %f51, %f47;
	mul.f32 	%f107, %f106, %f32;
	mul.f32 	%f108, %f106, %f29;
	mul.f32 	%f109, %f106, %f27;
	mul.f32 	%f110, %f105, %f107;
	mul.f32 	%f111, %f105, %f108;
	mul.f32 	%f112, %f105, %f109;
	fma.rn.f32 	%f113, %f100, %f102, %f110;
	fma.rn.f32 	%f114, %f100, %f103, %f111;
	fma.rn.f32 	%f115, %f100, %f104, %f112;
	fma.rn.f32 	%f116, %f57, %f57, %f98;
	mul.f32 	%f117, %f52, %f50;
	mul.f32 	%f118, %f117, %f31;
	mul.f32 	%f119, %f117, %f35;
	mul.f32 	%f120, %f117, %f38;
	fma.rn.f32 	%f121, %f116, %f118, %f113;
	fma.rn.f32 	%f122, %f116, %f119, %f114;
	fma.rn.f32 	%f123, %f116, %f120, %f115;
	mul.f32 	%f124, %f3, 0f40800000;
	mul.f32 	%f125, %f124, %f121;
	mul.f32 	%f126, %f124, %f122;
	mul.f32 	%f127, %f124, %f123;
	sub.f32 	%f128, %f95, %f125;
	sub.f32 	%f129, %f96, %f126;
	sub.f32 	%f130, %f97, %f127;
	cvta.to.global.u64 	%rd48, %rd2;
	add.s64 	%rd49, %rd48, %rd42;
	ld.global.f32 	%f131, [%rd49];
	add.f32 	%f132, %f131, %f128;
	st.global.f32 	[%rd49], %f132;
	cvta.to.global.u64 	%rd50, %rd3;
	add.s64 	%rd51, %rd50, %rd42;
	ld.global.f32 	%f133, [%rd51];
	add.f32 	%f134, %f133, %f129;
	st.global.f32 	[%rd51], %f134;
	cvta.to.global.u64 	%rd52, %rd4;
	add.s64 	%rd53, %rd52, %rd42;
	ld.global.f32 	%f135, [%rd53];
	add.f32 	%f136, %f135, %f130;
	st.global.f32 	[%rd53], %f136;

BB0_8:
	ret;
}


`
	addcubicanisotropy_ptx_35 = `
.version 4.0
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry addcubicanisotropy(
	.param .u64 addcubicanisotropy_param_0,
	.param .u64 addcubicanisotropy_param_1,
	.param .u64 addcubicanisotropy_param_2,
	.param .u64 addcubicanisotropy_param_3,
	.param .u64 addcubicanisotropy_param_4,
	.param .u64 addcubicanisotropy_param_5,
	.param .u64 addcubicanisotropy_param_6,
	.param .u64 addcubicanisotropy_param_7,
	.param .u64 addcubicanisotropy_param_8,
	.param .u64 addcubicanisotropy_param_9,
	.param .u64 addcubicanisotropy_param_10,
	.param .u64 addcubicanisotropy_param_11,
	.param .u64 addcubicanisotropy_param_12,
	.param .u64 addcubicanisotropy_param_13,
	.param .u64 addcubicanisotropy_param_14,
	.param .u64 addcubicanisotropy_param_15,
	.param .u32 addcubicanisotropy_param_16
)
{
	.reg .pred 	%p<4>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<139>;
	.reg .s64 	%rd<55>;


	ld.param.u64 	%rd3, [addcubicanisotropy_param_0];
	ld.param.u64 	%rd4, [addcubicanisotropy_param_1];
	ld.param.u64 	%rd5, [addcubicanisotropy_param_2];
	ld.param.u64 	%rd6, [addcubicanisotropy_param_3];
	ld.param.u64 	%rd7, [addcubicanisotropy_param_4];
	ld.param.u64 	%rd8, [addcubicanisotropy_param_5];
	ld.param.u64 	%rd9, [addcubicanisotropy_param_6];
	ld.param.u64 	%rd10, [addcubicanisotropy_param_7];
	ld.param.u64 	%rd11, [addcubicanisotropy_param_8];
	ld.param.u64 	%rd12, [addcubicanisotropy_param_9];
	ld.param.u64 	%rd13, [addcubicanisotropy_param_10];
	ld.param.u64 	%rd14, [addcubicanisotropy_param_11];
	ld.param.u64 	%rd15, [addcubicanisotropy_param_12];
	ld.param.u64 	%rd16, [addcubicanisotropy_param_13];
	ld.param.u64 	%rd17, [addcubicanisotropy_param_14];
	ld.param.u64 	%rd18, [addcubicanisotropy_param_15];
	ld.param.u32 	%r2, [addcubicanisotropy_param_16];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB2_8;

	cvta.to.global.u64 	%rd19, %rd14;
	cvta.to.global.u64 	%rd20, %rd13;
	cvta.to.global.u64 	%rd21, %rd12;
	cvta.to.global.u64 	%rd22, %rd11;
	cvta.to.global.u64 	%rd23, %rd10;
	cvta.to.global.u64 	%rd24, %rd9;
	cvta.to.global.u64 	%rd25, %rd18;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd26, %rd25, %rd1;
	ld.global.nc.u8 	%rs1, [%rd26];
	cvt.u64.u16	%rd27, %rs1;
	and.b64  	%rd2, %rd27, 255;
	shl.b64 	%rd28, %rd2, 2;
	add.s64 	%rd29, %rd24, %rd28;
	ld.global.nc.f32 	%f1, [%rd29];
	add.s64 	%rd30, %rd23, %rd28;
	ld.global.nc.f32 	%f2, [%rd30];
	add.s64 	%rd31, %rd22, %rd28;
	ld.global.nc.f32 	%f3, [%rd31];
	add.s64 	%rd32, %rd21, %rd28;
	add.s64 	%rd33, %rd20, %rd28;
	add.s64 	%rd34, %rd19, %rd28;
	ld.global.nc.f32 	%f4, [%rd32];
	ld.global.nc.f32 	%f5, [%rd33];
	mul.f32 	%f16, %f5, %f5;
	fma.rn.f32 	%f17, %f4, %f4, %f16;
	ld.global.nc.f32 	%f6, [%rd34];
	fma.rn.f32 	%f18, %f6, %f6, %f17;
	sqrt.rn.f32 	%f7, %f18;
	setp.neu.f32	%p2, %f7, 0f00000000;
	@%p2 bra 	BB2_3;

	mov.f32 	%f137, 0f00000000;
	bra.uni 	BB2_4;

BB2_3:
	rcp.rn.f32 	%f137, %f7;

BB2_4:
	cvta.to.global.u64 	%rd35, %rd17;
	cvta.to.global.u64 	%rd36, %rd16;
	cvta.to.global.u64 	%rd37, %rd15;
	add.s64 	%rd39, %rd37, %rd28;
	ld.global.nc.f32 	%f10, [%rd39];
	add.s64 	%rd40, %rd36, %rd28;
	ld.global.nc.f32 	%f11, [%rd40];
	mul.f32 	%f20, %f11, %f11;
	fma.rn.f32 	%f21, %f10, %f10, %f20;
	add.s64 	%rd41, %rd35, %rd28;
	ld.global.nc.f32 	%f12, [%rd41];
	fma.rn.f32 	%f22, %f12, %f12, %f21;
	sqrt.rn.f32 	%f13, %f22;
	setp.neu.f32	%p3, %f13, 0f00000000;
	@%p3 bra 	BB2_6;

	mov.f32 	%f138, 0f00000000;
	bra.uni 	BB2_7;

BB2_6:
	rcp.rn.f32 	%f138, %f13;

BB2_7:
	cvta.to.global.u64 	%rd42, %rd5;
	cvta.to.global.u64 	%rd43, %rd4;
	cvta.to.global.u64 	%rd44, %rd3;
	cvta.to.global.u64 	%rd45, %rd8;
	cvta.to.global.u64 	%rd46, %rd7;
	cvta.to.global.u64 	%rd47, %rd6;
	mul.f32 	%f24, %f137, %f4;
	mul.f32 	%f25, %f137, %f5;
	mul.f32 	%f26, %f137, %f6;
	mul.f32 	%f27, %f138, %f12;
	mul.f32 	%f28, %f25, %f27;
	mul.f32 	%f29, %f138, %f11;
	mul.f32 	%f30, %f26, %f29;
	sub.f32 	%f31, %f28, %f30;
	mul.f32 	%f32, %f138, %f10;
	mul.f32 	%f33, %f26, %f32;
	mul.f32 	%f34, %f24, %f27;
	sub.f32 	%f35, %f33, %f34;
	mul.f32 	%f36, %f24, %f29;
	mul.f32 	%f37, %f25, %f32;
	sub.f32 	%f38, %f36, %f37;
	shl.b64 	%rd48, %rd1, 2;
	add.s64 	%rd49, %rd47, %rd48;
	ld.global.nc.f32 	%f39, [%rd49];
	add.s64 	%rd50, %rd46, %rd48;
	ld.global.nc.f32 	%f40, [%rd50];
	mul.f32 	%f41, %f25, %f40;
	fma.rn.f32 	%f42, %f24, %f39, %f41;
	add.s64 	%rd51, %rd45, %rd48;
	ld.global.nc.f32 	%f43, [%rd51];
	fma.rn.f32 	%f44, %f26, %f43, %f42;
	mul.f32 	%f45, %f29, %f40;
	fma.rn.f32 	%f46, %f32, %f39, %f45;
	fma.rn.f32 	%f47, %f27, %f43, %f46;
	mul.f32 	%f48, %f35, %f40;
	fma.rn.f32 	%f49, %f31, %f39, %f48;
	fma.rn.f32 	%f50, %f38, %f43, %f49;
	mul.f32 	%f51, %f47, %f47;
	mul.f32 	%f52, %f50, %f50;
	add.f32 	%f53, %f51, %f52;
	mul.f32 	%f54, %f44, %f24;
	mul.f32 	%f55, %f44, %f25;
	mul.f32 	%f56, %f44, %f26;
	mul.f32 	%f57, %f44, %f44;
	add.f32 	%f58, %f57, %f52;
	mul.f32 	%f59, %f47, %f32;
	mul.f32 	%f60, %f47, %f29;
	mul.f32 	%f61, %f47, %f27;
	mul.f32 	%f62, %f58, %f59;
	mul.f32 	%f63, %f58, %f60;
	mul.f32 	%f64, %f58, %f61;
	fma.rn.f32 	%f65, %f53, %f54, %f62;
	fma.rn.f32 	%f66, %f53, %f55, %f63;
	fma.rn.f32 	%f67, %f53, %f56, %f64;
	add.f32 	%f68, %f57, %f51;
	mul.f32 	%f69, %f50, %f31;
	mul.f32 	%f70, %f50, %f35;
	mul.f32 	%f71, %f50, %f38;
	fma.rn.f32 	%f72, %f68, %f69, %f65;
	fma.rn.f32 	%f73, %f68, %f70, %f66;
	fma.rn.f32 	%f74, %f68, %f71, %f67;
	mul.f32 	%f75, %f1, 0fC0000000;
	mul.f32 	%f76, %f75, %f72;
	mul.f32 	%f77, %f75, %f73;
	mul.f32 	%f78, %f75, %f74;
	mul.f32 	%f79, %f51, %f52;
	mul.f32 	%f80, %f57, %f52;
	mul.f32 	%f81, %f80, %f59;
	mul.f32 	%f82, %f80, %f60;
	mul.f32 	%f83, %f80, %f61;
	fma.rn.f32 	%f84, %f79, %f54, %f81;
	fma.rn.f32 	%f85, %f79, %f55, %f82;
	fma.rn.f32 	%f86, %f79, %f56, %f83;
	mul.f32 	%f87, %f57, %f51;
	fma.rn.f32 	%f88, %f87, %f69, %f84;
	fma.rn.f32 	%f89, %f87, %f70, %f85;
	fma.rn.f32 	%f90, %f87, %f71, %f86;
	add.f32 	%f91, %f2, %f2;
	mul.f32 	%f92, %f91, %f88;
	mul.f32 	%f93, %f91, %f89;
	mul.f32 	%f94, %f91, %f90;
	sub.f32 	%f95, %f76, %f92;
	sub.f32 	%f96, %f77, %f93;
	sub.f32 	%f97, %f78, %f94;
	mul.f32 	%f98, %f51, %f51;
	mul.f32 	%f99, %f52, %f52;
	fma.rn.f32 	%f100, %f52, %f52, %f98;
	mul.f32 	%f101, %f57, %f44;
	mul.f32 	%f102, %f101, %f24;
	mul.f32 	%f103, %f101, %f25;
	mul.f32 	%f104, %f101, %f26;
	fma.rn.f32 	%f105, %f57, %f57, %f99;
	mul.f32 	%f106, %f51, %f47;
	mul.f32 	%f107, %f106, %f32;
	mul.f32 	%f108, %f106, %f29;
	mul.f32 	%f109, %f106, %f27;
	mul.f32 	%f110, %f105, %f107;
	mul.f32 	%f111, %f105, %f108;
	mul.f32 	%f112, %f105, %f109;
	fma.rn.f32 	%f113, %f100, %f102, %f110;
	fma.rn.f32 	%f114, %f100, %f103, %f111;
	fma.rn.f32 	%f115, %f100, %f104, %f112;
	fma.rn.f32 	%f116, %f57, %f57, %f98;
	mul.f32 	%f117, %f52, %f50;
	mul.f32 	%f118, %f117, %f31;
	mul.f32 	%f119, %f117, %f35;
	mul.f32 	%f120, %f117, %f38;
	fma.rn.f32 	%f121, %f116, %f118, %f113;
	fma.rn.f32 	%f122, %f116, %f119, %f114;
	fma.rn.f32 	%f123, %f116, %f120, %f115;
	mul.f32 	%f124, %f3, 0f40800000;
	mul.f32 	%f125, %f124, %f121;
	mul.f32 	%f126, %f124, %f122;
	mul.f32 	%f127, %f124, %f123;
	sub.f32 	%f128, %f95, %f125;
	sub.f32 	%f129, %f96, %f126;
	sub.f32 	%f130, %f97, %f127;
	add.s64 	%rd52, %rd44, %rd48;
	ld.global.f32 	%f131, [%rd52];
	add.f32 	%f132, %f131, %f128;
	st.global.f32 	[%rd52], %f132;
	add.s64 	%rd53, %rd43, %rd48;
	ld.global.f32 	%f133, [%rd53];
	add.f32 	%f134, %f133, %f129;
	st.global.f32 	[%rd53], %f134;
	add.s64 	%rd54, %rd42, %rd48;
	ld.global.f32 	%f135, [%rd54];
	add.f32 	%f136, %f135, %f130;
	st.global.f32 	[%rd54], %f136;

BB2_8:
	ret;
}


`
)
