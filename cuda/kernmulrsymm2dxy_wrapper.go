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

// CUDA handle for kernmulRSymm2Dxy kernel
var kernmulRSymm2Dxy_code cu.Function

// Stores the arguments for kernmulRSymm2Dxy kernel invocation
type kernmulRSymm2Dxy_args_t struct {
	arg_fftMx  unsafe.Pointer
	arg_fftMy  unsafe.Pointer
	arg_fftKxx unsafe.Pointer
	arg_fftKyy unsafe.Pointer
	arg_fftKxy unsafe.Pointer
	arg_Nx     int
	arg_Ny     int
	argptr     [7]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for kernmulRSymm2Dxy kernel invocation
var kernmulRSymm2Dxy_args kernmulRSymm2Dxy_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	kernmulRSymm2Dxy_args.argptr[0] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftMx)
	kernmulRSymm2Dxy_args.argptr[1] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftMy)
	kernmulRSymm2Dxy_args.argptr[2] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKxx)
	kernmulRSymm2Dxy_args.argptr[3] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKyy)
	kernmulRSymm2Dxy_args.argptr[4] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKxy)
	kernmulRSymm2Dxy_args.argptr[5] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_Nx)
	kernmulRSymm2Dxy_args.argptr[6] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_Ny)
}

// Wrapper for kernmulRSymm2Dxy CUDA kernel, asynchronous.
func k_kernmulRSymm2Dxy_async(fftMx unsafe.Pointer, fftMy unsafe.Pointer, fftKxx unsafe.Pointer, fftKyy unsafe.Pointer, fftKxy unsafe.Pointer, Nx int, Ny int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	kernmulRSymm2Dxy_args.Lock()
	defer kernmulRSymm2Dxy_args.Unlock()

	if kernmulRSymm2Dxy_code == 0 {
		kernmulRSymm2Dxy_code = fatbinLoad(kernmulRSymm2Dxy_map, "kernmulRSymm2Dxy")
	}

	kernmulRSymm2Dxy_args.arg_fftMx = fftMx
	kernmulRSymm2Dxy_args.arg_fftMy = fftMy
	kernmulRSymm2Dxy_args.arg_fftKxx = fftKxx
	kernmulRSymm2Dxy_args.arg_fftKyy = fftKyy
	kernmulRSymm2Dxy_args.arg_fftKxy = fftKxy
	kernmulRSymm2Dxy_args.arg_Nx = Nx
	kernmulRSymm2Dxy_args.arg_Ny = Ny

	args := kernmulRSymm2Dxy_args.argptr[:]
	cu.LaunchKernel(kernmulRSymm2Dxy_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for kernmulRSymm2Dxy kernel.
var kernmulRSymm2Dxy_map = map[int]string{0: "",
	20: kernmulRSymm2Dxy_ptx_20,
	30: kernmulRSymm2Dxy_ptx_30,
	35: kernmulRSymm2Dxy_ptx_35}

// kernmulRSymm2Dxy PTX code for various compute capabilities.
const (
	kernmulRSymm2Dxy_ptx_20 = `
.version 4.0
.target sm_20
.address_size 64


.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<19>;
	.reg .f32 	%f<18>;
	.reg .s64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p2, %p1;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd5;
	cvta.to.global.u64 	%rd7, %rd4;
	cvta.to.global.u64 	%rd8, %rd3;
	cvta.to.global.u64 	%rd9, %rd2;
	cvta.to.global.u64 	%rd10, %rd1;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd11, %r12, 4;
	add.s64 	%rd12, %rd10, %rd11;
	ld.global.f32 	%f1, [%rd12+4];
	add.s64 	%rd13, %rd9, %rd11;
	ld.global.f32 	%f2, [%rd13+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd14, %r18, 4;
	add.s64 	%rd15, %rd8, %rd14;
	add.s64 	%rd16, %rd7, %rd14;
	ld.global.f32 	%f4, [%rd16];
	add.s64 	%rd17, %rd6, %rd14;
	ld.global.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.f32 	%f7, [%rd15];
	ld.global.f32 	%f8, [%rd12];
	ld.global.f32 	%f9, [%rd13];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd12], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd12+4], %f13;
	mul.f32 	%f14, %f9, %f4;
	fma.rn.f32 	%f15, %f8, %f6, %f14;
	st.global.f32 	[%rd13], %f15;
	mul.f32 	%f16, %f2, %f4;
	fma.rn.f32 	%f17, %f1, %f6, %f16;
	st.global.f32 	[%rd13+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_30 = `
.version 4.0
.target sm_30
.address_size 64


.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<19>;
	.reg .f32 	%f<18>;
	.reg .s64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p2, %p1;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd5;
	cvta.to.global.u64 	%rd7, %rd4;
	cvta.to.global.u64 	%rd8, %rd3;
	cvta.to.global.u64 	%rd9, %rd2;
	cvta.to.global.u64 	%rd10, %rd1;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd11, %r12, 4;
	add.s64 	%rd12, %rd10, %rd11;
	ld.global.f32 	%f1, [%rd12+4];
	add.s64 	%rd13, %rd9, %rd11;
	ld.global.f32 	%f2, [%rd13+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd14, %r18, 4;
	add.s64 	%rd15, %rd8, %rd14;
	add.s64 	%rd16, %rd7, %rd14;
	ld.global.f32 	%f4, [%rd16];
	add.s64 	%rd17, %rd6, %rd14;
	ld.global.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.f32 	%f7, [%rd15];
	ld.global.f32 	%f8, [%rd12];
	ld.global.f32 	%f9, [%rd13];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd12], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd12+4], %f13;
	mul.f32 	%f14, %f9, %f4;
	fma.rn.f32 	%f15, %f8, %f6, %f14;
	st.global.f32 	[%rd13], %f15;
	mul.f32 	%f16, %f2, %f4;
	fma.rn.f32 	%f17, %f1, %f6, %f16;
	st.global.f32 	[%rd13+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_35 = `
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

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<19>;
	.reg .f32 	%f<18>;
	.reg .s64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p2, %p1;
	@%p3 bra 	BB2_2;

	cvta.to.global.u64 	%rd6, %rd5;
	cvta.to.global.u64 	%rd7, %rd4;
	cvta.to.global.u64 	%rd8, %rd3;
	cvta.to.global.u64 	%rd9, %rd2;
	cvta.to.global.u64 	%rd10, %rd1;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd11, %r12, 4;
	add.s64 	%rd12, %rd10, %rd11;
	ld.global.f32 	%f1, [%rd12+4];
	add.s64 	%rd13, %rd9, %rd11;
	ld.global.f32 	%f2, [%rd13+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd14, %r18, 4;
	add.s64 	%rd15, %rd8, %rd14;
	add.s64 	%rd16, %rd7, %rd14;
	ld.global.nc.f32 	%f4, [%rd16];
	add.s64 	%rd17, %rd6, %rd14;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd15];
	ld.global.f32 	%f8, [%rd12];
	ld.global.f32 	%f9, [%rd13];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd12], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd12+4], %f13;
	mul.f32 	%f14, %f9, %f4;
	fma.rn.f32 	%f15, %f8, %f6, %f14;
	st.global.f32 	[%rd13], %f15;
	mul.f32 	%f16, %f2, %f4;
	fma.rn.f32 	%f17, %f1, %f6, %f16;
	st.global.f32 	[%rd13+4], %f17;

BB2_2:
	ret;
}


`
)
