package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var addexchangemask_code cu.Function

type addexchangemask_args struct {
	arg_Beff  unsafe.Pointer
	arg_m     unsafe.Pointer
	arg_maskX unsafe.Pointer
	arg_maskY unsafe.Pointer
	arg_maskZ unsafe.Pointer
	arg_wx    float32
	arg_wy    float32
	arg_wz    float32
	arg_N0    int
	arg_N1    int
	arg_N2    int
	argptr    [11]unsafe.Pointer
}

// Wrapper for addexchangemask CUDA kernel, asynchronous.
func k_addexchangemask_async(Beff unsafe.Pointer, m unsafe.Pointer, maskX unsafe.Pointer, maskY unsafe.Pointer, maskZ unsafe.Pointer, wx float32, wy float32, wz float32, N0 int, N1 int, N2 int, cfg *config, str cu.Stream) {
	if addexchangemask_code == 0 {
		addexchangemask_code = fatbinLoad(addexchangemask_map, "addexchangemask")
	}

	var a addexchangemask_args

	a.arg_Beff = Beff
	a.argptr[0] = unsafe.Pointer(&a.arg_Beff)
	a.arg_m = m
	a.argptr[1] = unsafe.Pointer(&a.arg_m)
	a.arg_maskX = maskX
	a.argptr[2] = unsafe.Pointer(&a.arg_maskX)
	a.arg_maskY = maskY
	a.argptr[3] = unsafe.Pointer(&a.arg_maskY)
	a.arg_maskZ = maskZ
	a.argptr[4] = unsafe.Pointer(&a.arg_maskZ)
	a.arg_wx = wx
	a.argptr[5] = unsafe.Pointer(&a.arg_wx)
	a.arg_wy = wy
	a.argptr[6] = unsafe.Pointer(&a.arg_wy)
	a.arg_wz = wz
	a.argptr[7] = unsafe.Pointer(&a.arg_wz)
	a.arg_N0 = N0
	a.argptr[8] = unsafe.Pointer(&a.arg_N0)
	a.arg_N1 = N1
	a.argptr[9] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[10] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	cu.LaunchKernel(addexchangemask_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for addexchangemask CUDA kernel, synchronized.
func k_addexchangemask(Beff unsafe.Pointer, m unsafe.Pointer, maskX unsafe.Pointer, maskY unsafe.Pointer, maskZ unsafe.Pointer, wx float32, wy float32, wz float32, N0 int, N1 int, N2 int, cfg *config) {
	str := stream()
	k_addexchangemask_async(Beff, m, maskX, maskY, maskZ, wx, wy, wz, N0, N1, N2, cfg, str)
	syncAndRecycle(str)
}

var addexchangemask_map = map[int]string{0: "",
	20: addexchangemask_ptx_20,
	30: addexchangemask_ptx_30,
	35: addexchangemask_ptx_35}

const (
	addexchangemask_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry addexchangemask(
	.param .u64 addexchangemask_param_0,
	.param .u64 addexchangemask_param_1,
	.param .u64 addexchangemask_param_2,
	.param .u64 addexchangemask_param_3,
	.param .u64 addexchangemask_param_4,
	.param .f32 addexchangemask_param_5,
	.param .f32 addexchangemask_param_6,
	.param .f32 addexchangemask_param_7,
	.param .u32 addexchangemask_param_8,
	.param .u32 addexchangemask_param_9,
	.param .u32 addexchangemask_param_10
)
{
	.reg .pred 	%p<14>;
	.reg .s32 	%r<146>;
	.reg .f32 	%f<57>;
	.reg .s64 	%rd<44>;


	ld.param.u64 	%rd3, [addexchangemask_param_0];
	ld.param.u64 	%rd4, [addexchangemask_param_1];
	ld.param.u64 	%rd5, [addexchangemask_param_2];
	ld.param.u64 	%rd6, [addexchangemask_param_3];
	ld.param.u64 	%rd7, [addexchangemask_param_4];
	ld.param.f32 	%f25, [addexchangemask_param_5];
	ld.param.f32 	%f26, [addexchangemask_param_6];
	ld.param.f32 	%f27, [addexchangemask_param_7];
	ld.param.u32 	%r24, [addexchangemask_param_8];
	ld.param.u32 	%r25, [addexchangemask_param_9];
	ld.param.u32 	%r26, [addexchangemask_param_10];
	.loc 2 12 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 2 13 1
	mov.u32 	%r27, %ntid.y;
	mov.u32 	%r28, %ctaid.y;
	mov.u32 	%r29, %tid.y;
	mad.lo.s32 	%r5, %r27, %r28, %r29;
	.loc 2 15 1
	setp.lt.s32 	%p1, %r5, %r26;
	setp.lt.s32 	%p2, %r4, %r25;
	and.pred  	%p3, %p2, %p1;
	.loc 2 19 1
	setp.gt.s32 	%p4, %r24, 0;
	.loc 2 15 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_17;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 25 1
	add.s32 	%r31, %r5, -1;
	mov.u32 	%r145, 0;
	.loc 3 238 5
	max.s32 	%r32, %r31, %r145;
	.loc 2 26 1
	add.s32 	%r33, %r26, -1;
	add.s32 	%r34, %r5, 1;
	.loc 3 210 5
	min.s32 	%r35, %r34, %r33;
	.loc 2 31 1
	add.s32 	%r36, %r4, -1;
	.loc 3 238 5
	max.s32 	%r37, %r36, %r145;
	.loc 2 32 1
	add.s32 	%r38, %r25, -1;
	add.s32 	%r39, %r4, 1;
	.loc 3 210 5
	min.s32 	%r40, %r39, %r38;
	.loc 2 19 1
	mad.lo.s32 	%r144, %r40, %r26, %r5;
	mad.lo.s32 	%r143, %r37, %r26, %r5;
	mad.lo.s32 	%r142, %r26, %r4, %r35;
	mad.lo.s32 	%r141, %r26, %r4, %r32;
	mad.lo.s32 	%r140, %r26, %r4, %r5;
	cvta.to.global.u64 	%rd8, %rd3;

BB0_2:
	.loc 2 22 1
	mov.u32 	%r16, %r145;
	cvt.s64.s32 	%rd1, %r140;
	mul.wide.s32 	%rd9, %r140, 4;
	add.s64 	%rd2, %rd8, %rd9;
	ld.global.f32 	%f1, [%rd2];
	cvta.to.global.u64 	%rd10, %rd4;
	.loc 2 23 1
	add.s64 	%rd11, %rd10, %rd9;
	ld.global.f32 	%f2, [%rd11];
	.loc 2 25 1
	mul.wide.s32 	%rd12, %r141, 4;
	add.s64 	%rd13, %rd10, %rd12;
	ld.global.f32 	%f3, [%rd13];
	.loc 2 26 1
	mul.wide.s32 	%rd14, %r142, 4;
	add.s64 	%rd15, %rd10, %rd14;
	ld.global.f32 	%f4, [%rd15];
	setp.eq.s64 	%p6, %rd5, 0;
	mov.f32 	%f28, 0f3F800000;
	.loc 2 27 1
	mov.f32 	%f52, %f28;
	@%p6 bra 	BB0_4;

	cvta.to.global.u64 	%rd16, %rd5;
	.loc 2 27 1
	shl.b64 	%rd17, %rd1, 2;
	add.s64 	%rd18, %rd16, %rd17;
	ld.global.f32 	%f5, [%rd18];
	mov.f32 	%f52, %f5;

BB0_4:
	.loc 2 27 1
	mov.f32 	%f6, %f52;
	.loc 2 28 1
	mov.f32 	%f51, %f28;
	@%p6 bra 	BB0_6;

	mov.u32 	%r51, 0;
	.loc 3 238 5
	max.s32 	%r52, %r16, %r51;
	.loc 2 39 1
	add.s32 	%r53, %r24, -1;
	.loc 3 210 5
	min.s32 	%r54, %r52, %r53;
	.loc 3 238 5
	max.s32 	%r59, %r4, %r51;
	.loc 3 210 5
	min.s32 	%r61, %r59, %r38;
	.loc 2 28 1
	mad.lo.s32 	%r62, %r54, %r25, %r61;
	.loc 3 238 5
	max.s32 	%r68, %r34, %r51;
	.loc 3 210 5
	min.s32 	%r70, %r68, %r33;
	.loc 2 28 1
	mad.lo.s32 	%r71, %r62, %r26, %r70;
	cvta.to.global.u64 	%rd19, %rd5;
	.loc 2 28 1
	mul.wide.s32 	%rd20, %r71, 4;
	add.s64 	%rd21, %rd19, %rd20;
	ld.global.f32 	%f51, [%rd21];

BB0_6:
	.loc 2 29 1
	sub.f32 	%f31, %f3, %f2;
	sub.f32 	%f32, %f4, %f2;
	mul.f32 	%f33, %f51, %f32;
	fma.rn.f32 	%f34, %f6, %f31, %f33;
	fma.rn.f32 	%f9, %f34, %f27, %f1;
	.loc 2 31 1
	mul.wide.s32 	%rd23, %r143, 4;
	add.s64 	%rd24, %rd10, %rd23;
	ld.global.f32 	%f10, [%rd24];
	.loc 2 32 1
	mul.wide.s32 	%rd25, %r144, 4;
	add.s64 	%rd26, %rd10, %rd25;
	ld.global.f32 	%f11, [%rd26];
	setp.eq.s64 	%p8, %rd6, 0;
	.loc 2 33 1
	mov.f32 	%f50, %f28;
	@%p8 bra 	BB0_8;

	cvta.to.global.u64 	%rd27, %rd6;
	.loc 2 33 1
	shl.b64 	%rd28, %rd1, 2;
	add.s64 	%rd29, %rd27, %rd28;
	ld.global.f32 	%f50, [%rd29];

BB0_8:
	.loc 2 34 1
	mov.f32 	%f49, %f28;
	@%p8 bra 	BB0_10;

	mov.u32 	%r76, 0;
	.loc 3 238 5
	max.s32 	%r77, %r16, %r76;
	.loc 2 39 1
	add.s32 	%r78, %r24, -1;
	.loc 3 210 5
	min.s32 	%r79, %r77, %r78;
	.loc 3 238 5
	max.s32 	%r85, %r39, %r76;
	.loc 3 210 5
	min.s32 	%r87, %r85, %r38;
	.loc 2 34 1
	mad.lo.s32 	%r88, %r79, %r25, %r87;
	.loc 3 238 5
	max.s32 	%r93, %r5, %r76;
	.loc 3 210 5
	min.s32 	%r95, %r93, %r33;
	.loc 2 34 1
	mad.lo.s32 	%r96, %r88, %r26, %r95;
	cvta.to.global.u64 	%rd30, %rd6;
	.loc 2 34 1
	mul.wide.s32 	%rd31, %r96, 4;
	add.s64 	%rd32, %rd30, %rd31;
	ld.global.f32 	%f49, [%rd32];

BB0_10:
	.loc 2 35 1
	sub.f32 	%f36, %f10, %f2;
	sub.f32 	%f37, %f11, %f2;
	mul.f32 	%f38, %f49, %f37;
	fma.rn.f32 	%f39, %f50, %f36, %f38;
	fma.rn.f32 	%f56, %f39, %f26, %f9;
	.loc 2 19 18
	add.s32 	%r145, %r16, 1;
	setp.eq.s32 	%p10, %r24, 1;
	.loc 2 38 1
	@%p10 bra 	BB0_16;

	setp.eq.s64 	%p11, %rd7, 0;
	.loc 2 39 1
	add.s32 	%r98, %r24, -1;
	.loc 3 210 5
	min.s32 	%r99, %r145, %r98;
	.loc 2 39 1
	mad.lo.s32 	%r104, %r99, %r25, %r4;
	mad.lo.s32 	%r109, %r104, %r26, %r5;
	mul.wide.s32 	%rd34, %r109, 4;
	add.s64 	%rd35, %rd10, %rd34;
	ld.global.f32 	%f17, [%rd35];
	.loc 2 40 1
	add.s32 	%r111, %r16, -1;
	mov.u32 	%r112, 0;
	.loc 3 238 5
	max.s32 	%r113, %r111, %r112;
	.loc 2 40 1
	mad.lo.s32 	%r114, %r113, %r25, %r4;
	mad.lo.s32 	%r115, %r114, %r26, %r5;
	mul.wide.s32 	%rd36, %r115, 4;
	add.s64 	%rd37, %rd10, %rd36;
	ld.global.f32 	%f18, [%rd37];
	mov.f32 	%f40, 0f3F800000;
	.loc 2 41 1
	mov.f32 	%f55, %f40;
	@%p11 bra 	BB0_13;

	cvta.to.global.u64 	%rd38, %rd7;
	.loc 2 41 1
	shl.b64 	%rd39, %rd1, 2;
	add.s64 	%rd40, %rd38, %rd39;
	ld.global.f32 	%f19, [%rd40];
	mov.f32 	%f55, %f19;

BB0_13:
	.loc 2 41 1
	mov.f32 	%f20, %f55;
	.loc 2 42 1
	mov.f32 	%f54, %f40;
	@%p11 bra 	BB0_15;

	.loc 3 238 5
	max.s32 	%r119, %r145, %r112;
	.loc 3 210 5
	min.s32 	%r121, %r119, %r98;
	.loc 3 238 5
	max.s32 	%r126, %r4, %r112;
	.loc 3 210 5
	min.s32 	%r128, %r126, %r38;
	.loc 2 42 1
	mad.lo.s32 	%r129, %r121, %r25, %r128;
	.loc 3 238 5
	max.s32 	%r134, %r5, %r112;
	.loc 3 210 5
	min.s32 	%r136, %r134, %r33;
	.loc 2 42 1
	mad.lo.s32 	%r137, %r129, %r26, %r136;
	cvta.to.global.u64 	%rd41, %rd7;
	.loc 2 42 1
	mul.wide.s32 	%rd42, %r137, 4;
	add.s64 	%rd43, %rd41, %rd42;
	ld.global.f32 	%f54, [%rd43];

BB0_15:
	.loc 2 43 1
	sub.f32 	%f42, %f17, %f2;
	sub.f32 	%f43, %f18, %f2;
	mul.f32 	%f44, %f54, %f43;
	fma.rn.f32 	%f45, %f20, %f42, %f44;
	fma.rn.f32 	%f56, %f45, %f25, %f56;

BB0_16:
	.loc 2 46 1
	st.global.f32 	[%rd2], %f56;
	.loc 2 19 1
	mad.lo.s32 	%r144, %r26, %r25, %r144;
	mad.lo.s32 	%r143, %r26, %r25, %r143;
	mad.lo.s32 	%r142, %r26, %r25, %r142;
	mad.lo.s32 	%r141, %r26, %r25, %r141;
	mad.lo.s32 	%r140, %r26, %r25, %r140;
	setp.lt.s32 	%p13, %r145, %r24;
	@%p13 bra 	BB0_2;

BB0_17:
	.loc 2 48 2
	ret;
}


`
	addexchangemask_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry addexchangemask(
	.param .u64 addexchangemask_param_0,
	.param .u64 addexchangemask_param_1,
	.param .u64 addexchangemask_param_2,
	.param .u64 addexchangemask_param_3,
	.param .u64 addexchangemask_param_4,
	.param .f32 addexchangemask_param_5,
	.param .f32 addexchangemask_param_6,
	.param .f32 addexchangemask_param_7,
	.param .u32 addexchangemask_param_8,
	.param .u32 addexchangemask_param_9,
	.param .u32 addexchangemask_param_10
)
{
	.reg .pred 	%p<14>;
	.reg .s32 	%r<112>;
	.reg .f32 	%f<57>;
	.reg .s64 	%rd<44>;


	ld.param.u64 	%rd3, [addexchangemask_param_0];
	ld.param.u64 	%rd4, [addexchangemask_param_1];
	ld.param.u64 	%rd5, [addexchangemask_param_2];
	ld.param.u64 	%rd6, [addexchangemask_param_3];
	ld.param.u64 	%rd7, [addexchangemask_param_4];
	ld.param.f32 	%f25, [addexchangemask_param_5];
	ld.param.f32 	%f26, [addexchangemask_param_6];
	ld.param.f32 	%f27, [addexchangemask_param_7];
	ld.param.u32 	%r28, [addexchangemask_param_8];
	ld.param.u32 	%r29, [addexchangemask_param_9];
	ld.param.u32 	%r30, [addexchangemask_param_10];
	.loc 2 12 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 2 13 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 15 1
	setp.lt.s32 	%p1, %r8, %r30;
	setp.lt.s32 	%p2, %r4, %r29;
	and.pred  	%p3, %p2, %p1;
	.loc 2 19 1
	setp.gt.s32 	%p4, %r28, 0;
	.loc 2 15 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_17;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 25 1
	add.s32 	%r32, %r8, -1;
	mov.u32 	%r111, 0;
	.loc 3 238 5
	max.s32 	%r33, %r32, %r111;
	.loc 2 26 1
	add.s32 	%r34, %r30, -1;
	add.s32 	%r35, %r8, 1;
	.loc 3 210 5
	min.s32 	%r36, %r35, %r34;
	.loc 2 31 1
	add.s32 	%r37, %r4, -1;
	.loc 3 238 5
	max.s32 	%r38, %r37, %r111;
	.loc 2 32 1
	add.s32 	%r39, %r29, -1;
	add.s32 	%r40, %r4, 1;
	.loc 3 210 5
	min.s32 	%r41, %r40, %r39;
	.loc 2 19 1
	mad.lo.s32 	%r110, %r41, %r30, %r8;
	mul.lo.s32 	%r10, %r30, %r29;
	mad.lo.s32 	%r109, %r38, %r30, %r8;
	mad.lo.s32 	%r108, %r30, %r4, %r36;
	mad.lo.s32 	%r107, %r30, %r4, %r33;
	mad.lo.s32 	%r106, %r30, %r4, %r8;
	cvta.to.global.u64 	%rd8, %rd3;

BB0_2:
	.loc 2 22 1
	mov.u32 	%r20, %r111;
	cvt.s64.s32 	%rd1, %r106;
	mul.wide.s32 	%rd9, %r106, 4;
	add.s64 	%rd2, %rd8, %rd9;
	ld.global.f32 	%f1, [%rd2];
	cvta.to.global.u64 	%rd10, %rd4;
	.loc 2 23 1
	add.s64 	%rd11, %rd10, %rd9;
	ld.global.f32 	%f2, [%rd11];
	.loc 2 25 1
	mul.wide.s32 	%rd12, %r107, 4;
	add.s64 	%rd13, %rd10, %rd12;
	ld.global.f32 	%f3, [%rd13];
	.loc 2 26 1
	mul.wide.s32 	%rd14, %r108, 4;
	add.s64 	%rd15, %rd10, %rd14;
	ld.global.f32 	%f4, [%rd15];
	setp.eq.s64 	%p6, %rd5, 0;
	mov.f32 	%f28, 0f3F800000;
	.loc 2 27 1
	mov.f32 	%f52, %f28;
	@%p6 bra 	BB0_4;

	cvta.to.global.u64 	%rd16, %rd5;
	.loc 2 27 1
	shl.b64 	%rd17, %rd1, 2;
	add.s64 	%rd18, %rd16, %rd17;
	ld.global.f32 	%f5, [%rd18];
	mov.f32 	%f52, %f5;

BB0_4:
	.loc 2 27 1
	mov.f32 	%f6, %f52;
	.loc 2 28 1
	mov.f32 	%f51, %f28;
	@%p6 bra 	BB0_6;

	mov.u32 	%r49, 0;
	.loc 3 238 5
	max.s32 	%r50, %r20, %r49;
	.loc 2 39 1
	add.s32 	%r51, %r28, -1;
	.loc 3 210 5
	min.s32 	%r52, %r50, %r51;
	.loc 3 238 5
	max.s32 	%r54, %r4, %r49;
	.loc 3 210 5
	min.s32 	%r55, %r54, %r39;
	.loc 2 28 1
	mad.lo.s32 	%r56, %r52, %r29, %r55;
	.loc 3 238 5
	max.s32 	%r58, %r35, %r49;
	.loc 3 210 5
	min.s32 	%r60, %r58, %r34;
	.loc 2 28 1
	mad.lo.s32 	%r61, %r56, %r30, %r60;
	cvta.to.global.u64 	%rd19, %rd5;
	.loc 2 28 1
	mul.wide.s32 	%rd20, %r61, 4;
	add.s64 	%rd21, %rd19, %rd20;
	ld.global.f32 	%f51, [%rd21];

BB0_6:
	.loc 2 29 1
	sub.f32 	%f31, %f3, %f2;
	sub.f32 	%f32, %f4, %f2;
	mul.f32 	%f33, %f51, %f32;
	fma.rn.f32 	%f34, %f6, %f31, %f33;
	fma.rn.f32 	%f9, %f34, %f27, %f1;
	.loc 2 31 1
	mul.wide.s32 	%rd23, %r109, 4;
	add.s64 	%rd24, %rd10, %rd23;
	ld.global.f32 	%f10, [%rd24];
	.loc 2 32 1
	mul.wide.s32 	%rd25, %r110, 4;
	add.s64 	%rd26, %rd10, %rd25;
	ld.global.f32 	%f11, [%rd26];
	setp.eq.s64 	%p8, %rd6, 0;
	.loc 2 33 1
	mov.f32 	%f50, %f28;
	@%p8 bra 	BB0_8;

	cvta.to.global.u64 	%rd27, %rd6;
	.loc 2 33 1
	shl.b64 	%rd28, %rd1, 2;
	add.s64 	%rd29, %rd27, %rd28;
	ld.global.f32 	%f50, [%rd29];

BB0_8:
	.loc 2 34 1
	mov.f32 	%f49, %f28;
	@%p8 bra 	BB0_10;

	mov.u32 	%r66, 0;
	.loc 3 238 5
	max.s32 	%r67, %r20, %r66;
	.loc 2 39 1
	add.s32 	%r68, %r28, -1;
	.loc 3 210 5
	min.s32 	%r69, %r67, %r68;
	.loc 3 238 5
	max.s32 	%r71, %r40, %r66;
	.loc 3 210 5
	min.s32 	%r73, %r71, %r39;
	.loc 2 34 1
	mad.lo.s32 	%r74, %r69, %r29, %r73;
	.loc 3 238 5
	max.s32 	%r76, %r8, %r66;
	.loc 3 210 5
	min.s32 	%r77, %r76, %r34;
	.loc 2 34 1
	mad.lo.s32 	%r78, %r74, %r30, %r77;
	cvta.to.global.u64 	%rd30, %rd6;
	.loc 2 34 1
	mul.wide.s32 	%rd31, %r78, 4;
	add.s64 	%rd32, %rd30, %rd31;
	ld.global.f32 	%f49, [%rd32];

BB0_10:
	.loc 2 35 1
	sub.f32 	%f36, %f10, %f2;
	sub.f32 	%f37, %f11, %f2;
	mul.f32 	%f38, %f49, %f37;
	fma.rn.f32 	%f39, %f50, %f36, %f38;
	fma.rn.f32 	%f56, %f39, %f26, %f9;
	.loc 2 19 18
	add.s32 	%r111, %r20, 1;
	setp.eq.s32 	%p10, %r28, 1;
	.loc 2 38 1
	@%p10 bra 	BB0_16;

	setp.eq.s64 	%p11, %rd7, 0;
	.loc 2 39 1
	add.s32 	%r80, %r28, -1;
	.loc 3 210 5
	min.s32 	%r81, %r111, %r80;
	.loc 2 39 1
	mad.lo.s32 	%r82, %r81, %r29, %r4;
	mad.lo.s32 	%r83, %r82, %r30, %r8;
	mul.wide.s32 	%rd34, %r83, 4;
	add.s64 	%rd35, %rd10, %rd34;
	ld.global.f32 	%f17, [%rd35];
	.loc 2 40 1
	add.s32 	%r85, %r20, -1;
	mov.u32 	%r86, 0;
	.loc 3 238 5
	max.s32 	%r87, %r85, %r86;
	.loc 2 40 1
	mad.lo.s32 	%r88, %r87, %r29, %r4;
	mad.lo.s32 	%r89, %r88, %r30, %r8;
	mul.wide.s32 	%rd36, %r89, 4;
	add.s64 	%rd37, %rd10, %rd36;
	ld.global.f32 	%f18, [%rd37];
	mov.f32 	%f40, 0f3F800000;
	.loc 2 41 1
	mov.f32 	%f55, %f40;
	@%p11 bra 	BB0_13;

	cvta.to.global.u64 	%rd38, %rd7;
	.loc 2 41 1
	shl.b64 	%rd39, %rd1, 2;
	add.s64 	%rd40, %rd38, %rd39;
	ld.global.f32 	%f19, [%rd40];
	mov.f32 	%f55, %f19;

BB0_13:
	.loc 2 41 1
	mov.f32 	%f20, %f55;
	.loc 2 42 1
	mov.f32 	%f54, %f40;
	@%p11 bra 	BB0_15;

	.loc 3 238 5
	max.s32 	%r93, %r111, %r86;
	.loc 3 210 5
	min.s32 	%r95, %r93, %r80;
	.loc 3 238 5
	max.s32 	%r97, %r4, %r86;
	.loc 3 210 5
	min.s32 	%r98, %r97, %r39;
	.loc 2 42 1
	mad.lo.s32 	%r99, %r95, %r29, %r98;
	.loc 3 238 5
	max.s32 	%r101, %r8, %r86;
	.loc 3 210 5
	min.s32 	%r102, %r101, %r34;
	.loc 2 42 1
	mad.lo.s32 	%r103, %r99, %r30, %r102;
	cvta.to.global.u64 	%rd41, %rd7;
	.loc 2 42 1
	mul.wide.s32 	%rd42, %r103, 4;
	add.s64 	%rd43, %rd41, %rd42;
	ld.global.f32 	%f54, [%rd43];

BB0_15:
	.loc 2 43 1
	sub.f32 	%f42, %f17, %f2;
	sub.f32 	%f43, %f18, %f2;
	mul.f32 	%f44, %f54, %f43;
	fma.rn.f32 	%f45, %f20, %f42, %f44;
	fma.rn.f32 	%f56, %f45, %f25, %f56;

BB0_16:
	.loc 2 46 1
	st.global.f32 	[%rd2], %f56;
	.loc 2 19 1
	add.s32 	%r110, %r110, %r10;
	add.s32 	%r109, %r109, %r10;
	add.s32 	%r108, %r108, %r10;
	add.s32 	%r107, %r107, %r10;
	add.s32 	%r106, %r106, %r10;
	setp.lt.s32 	%p13, %r111, %r28;
	@%p13 bra 	BB0_2;

BB0_17:
	.loc 2 48 2
	ret;
}


`
	addexchangemask_ptx_35 = `
.version 3.1
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
	.loc 2 66 3
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
	.loc 2 71 3
	ret;
}

.visible .entry addexchangemask(
	.param .u64 addexchangemask_param_0,
	.param .u64 addexchangemask_param_1,
	.param .u64 addexchangemask_param_2,
	.param .u64 addexchangemask_param_3,
	.param .u64 addexchangemask_param_4,
	.param .f32 addexchangemask_param_5,
	.param .f32 addexchangemask_param_6,
	.param .f32 addexchangemask_param_7,
	.param .u32 addexchangemask_param_8,
	.param .u32 addexchangemask_param_9,
	.param .u32 addexchangemask_param_10
)
{
	.reg .pred 	%p<14>;
	.reg .s32 	%r<99>;
	.reg .f32 	%f<57>;
	.reg .s64 	%rd<44>;


	ld.param.u64 	%rd3, [addexchangemask_param_0];
	ld.param.u64 	%rd4, [addexchangemask_param_1];
	ld.param.u64 	%rd5, [addexchangemask_param_2];
	ld.param.u64 	%rd6, [addexchangemask_param_3];
	ld.param.u64 	%rd7, [addexchangemask_param_4];
	ld.param.f32 	%f25, [addexchangemask_param_5];
	ld.param.f32 	%f26, [addexchangemask_param_6];
	ld.param.f32 	%f27, [addexchangemask_param_7];
	ld.param.u32 	%r28, [addexchangemask_param_8];
	ld.param.u32 	%r29, [addexchangemask_param_9];
	ld.param.u32 	%r30, [addexchangemask_param_10];
	.loc 3 12 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 3 13 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 3 15 1
	setp.lt.s32 	%p1, %r8, %r30;
	setp.lt.s32 	%p2, %r4, %r29;
	and.pred  	%p3, %p2, %p1;
	.loc 3 19 1
	setp.gt.s32 	%p4, %r28, 0;
	.loc 3 15 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_17;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 25 1
	add.s32 	%r32, %r8, -1;
	mov.u32 	%r98, 0;
	.loc 4 238 5
	max.s32 	%r33, %r32, %r98;
	.loc 3 26 1
	add.s32 	%r34, %r30, -1;
	add.s32 	%r35, %r8, 1;
	.loc 4 210 5
	min.s32 	%r36, %r35, %r34;
	.loc 3 31 1
	add.s32 	%r37, %r4, -1;
	.loc 4 238 5
	max.s32 	%r38, %r37, %r98;
	.loc 3 32 1
	add.s32 	%r39, %r29, -1;
	add.s32 	%r40, %r4, 1;
	.loc 4 210 5
	min.s32 	%r41, %r40, %r39;
	.loc 3 19 1
	mad.lo.s32 	%r97, %r41, %r30, %r8;
	mul.lo.s32 	%r10, %r30, %r29;
	mad.lo.s32 	%r96, %r38, %r30, %r8;
	mad.lo.s32 	%r95, %r30, %r4, %r36;
	mad.lo.s32 	%r94, %r30, %r4, %r33;
	mad.lo.s32 	%r93, %r30, %r4, %r8;
	cvta.to.global.u64 	%rd8, %rd3;

BB2_2:
	.loc 3 22 1
	mov.u32 	%r20, %r98;
	cvt.s64.s32 	%rd1, %r93;
	mul.wide.s32 	%rd9, %r93, 4;
	add.s64 	%rd2, %rd8, %rd9;
	ld.global.f32 	%f1, [%rd2];
	cvta.to.global.u64 	%rd10, %rd4;
	.loc 3 23 1
	add.s64 	%rd11, %rd10, %rd9;
	ld.global.nc.f32 	%f2, [%rd11];
	.loc 3 25 1
	mul.wide.s32 	%rd12, %r94, 4;
	add.s64 	%rd13, %rd10, %rd12;
	ld.global.nc.f32 	%f3, [%rd13];
	.loc 3 26 1
	mul.wide.s32 	%rd14, %r95, 4;
	add.s64 	%rd15, %rd10, %rd14;
	ld.global.nc.f32 	%f4, [%rd15];
	setp.eq.s64 	%p6, %rd5, 0;
	mov.f32 	%f28, 0f3F800000;
	.loc 3 27 1
	mov.f32 	%f52, %f28;
	@%p6 bra 	BB2_4;

	cvta.to.global.u64 	%rd16, %rd5;
	.loc 3 27 1
	shl.b64 	%rd17, %rd1, 2;
	add.s64 	%rd18, %rd16, %rd17;
	ld.global.nc.f32 	%f5, [%rd18];
	mov.f32 	%f52, %f5;

BB2_4:
	.loc 3 27 1
	mov.f32 	%f6, %f52;
	.loc 3 28 1
	mov.f32 	%f51, %f28;
	@%p6 bra 	BB2_6;

	mov.u32 	%r45, 0;
	.loc 4 238 5
	max.s32 	%r46, %r20, %r45;
	.loc 3 39 1
	add.s32 	%r47, %r28, -1;
	.loc 4 210 5
	min.s32 	%r48, %r46, %r47;
	.loc 4 238 5
	max.s32 	%r50, %r4, %r45;
	.loc 4 210 5
	min.s32 	%r51, %r50, %r39;
	.loc 3 28 1
	mad.lo.s32 	%r52, %r48, %r29, %r51;
	.loc 4 238 5
	max.s32 	%r54, %r35, %r45;
	.loc 4 210 5
	min.s32 	%r56, %r54, %r34;
	.loc 3 28 1
	mad.lo.s32 	%r57, %r52, %r30, %r56;
	cvta.to.global.u64 	%rd19, %rd5;
	.loc 3 28 1
	mul.wide.s32 	%rd20, %r57, 4;
	add.s64 	%rd21, %rd19, %rd20;
	ld.global.nc.f32 	%f51, [%rd21];

BB2_6:
	.loc 3 29 1
	sub.f32 	%f31, %f3, %f2;
	sub.f32 	%f32, %f4, %f2;
	mul.f32 	%f33, %f51, %f32;
	fma.rn.f32 	%f34, %f6, %f31, %f33;
	fma.rn.f32 	%f9, %f34, %f27, %f1;
	.loc 3 31 1
	mul.wide.s32 	%rd23, %r96, 4;
	add.s64 	%rd24, %rd10, %rd23;
	ld.global.nc.f32 	%f10, [%rd24];
	.loc 3 32 1
	mul.wide.s32 	%rd25, %r97, 4;
	add.s64 	%rd26, %rd10, %rd25;
	ld.global.nc.f32 	%f11, [%rd26];
	setp.eq.s64 	%p8, %rd6, 0;
	.loc 3 33 1
	mov.f32 	%f50, %f28;
	@%p8 bra 	BB2_8;

	cvta.to.global.u64 	%rd27, %rd6;
	.loc 3 33 1
	shl.b64 	%rd28, %rd1, 2;
	add.s64 	%rd29, %rd27, %rd28;
	ld.global.nc.f32 	%f50, [%rd29];

BB2_8:
	.loc 3 34 1
	mov.f32 	%f49, %f28;
	@%p8 bra 	BB2_10;

	mov.u32 	%r58, 0;
	.loc 4 238 5
	max.s32 	%r59, %r20, %r58;
	.loc 3 39 1
	add.s32 	%r60, %r28, -1;
	.loc 4 210 5
	min.s32 	%r61, %r59, %r60;
	.loc 4 238 5
	max.s32 	%r63, %r40, %r58;
	.loc 4 210 5
	min.s32 	%r65, %r63, %r39;
	.loc 3 34 1
	mad.lo.s32 	%r66, %r61, %r29, %r65;
	.loc 4 238 5
	max.s32 	%r68, %r8, %r58;
	.loc 4 210 5
	min.s32 	%r69, %r68, %r34;
	.loc 3 34 1
	mad.lo.s32 	%r70, %r66, %r30, %r69;
	cvta.to.global.u64 	%rd30, %rd6;
	.loc 3 34 1
	mul.wide.s32 	%rd31, %r70, 4;
	add.s64 	%rd32, %rd30, %rd31;
	ld.global.nc.f32 	%f49, [%rd32];

BB2_10:
	.loc 3 35 1
	sub.f32 	%f36, %f10, %f2;
	sub.f32 	%f37, %f11, %f2;
	mul.f32 	%f38, %f49, %f37;
	fma.rn.f32 	%f39, %f50, %f36, %f38;
	fma.rn.f32 	%f56, %f39, %f26, %f9;
	.loc 3 19 18
	add.s32 	%r98, %r20, 1;
	setp.eq.s32 	%p10, %r28, 1;
	.loc 3 38 1
	@%p10 bra 	BB2_16;

	setp.eq.s64 	%p11, %rd7, 0;
	.loc 3 39 1
	add.s32 	%r71, %r28, -1;
	.loc 4 210 5
	min.s32 	%r72, %r98, %r71;
	.loc 3 39 1
	mad.lo.s32 	%r73, %r72, %r29, %r4;
	mad.lo.s32 	%r74, %r73, %r30, %r8;
	mul.wide.s32 	%rd34, %r74, 4;
	add.s64 	%rd35, %rd10, %rd34;
	ld.global.nc.f32 	%f17, [%rd35];
	.loc 3 40 1
	add.s32 	%r75, %r20, -1;
	mov.u32 	%r76, 0;
	.loc 4 238 5
	max.s32 	%r77, %r75, %r76;
	.loc 3 40 1
	mad.lo.s32 	%r78, %r77, %r29, %r4;
	mad.lo.s32 	%r79, %r78, %r30, %r8;
	mul.wide.s32 	%rd36, %r79, 4;
	add.s64 	%rd37, %rd10, %rd36;
	ld.global.nc.f32 	%f18, [%rd37];
	mov.f32 	%f40, 0f3F800000;
	.loc 3 41 1
	mov.f32 	%f55, %f40;
	@%p11 bra 	BB2_13;

	cvta.to.global.u64 	%rd38, %rd7;
	.loc 3 41 1
	shl.b64 	%rd39, %rd1, 2;
	add.s64 	%rd40, %rd38, %rd39;
	ld.global.nc.f32 	%f19, [%rd40];
	mov.f32 	%f55, %f19;

BB2_13:
	.loc 3 41 1
	mov.f32 	%f20, %f55;
	.loc 3 42 1
	mov.f32 	%f54, %f40;
	@%p11 bra 	BB2_15;

	.loc 4 238 5
	max.s32 	%r81, %r98, %r76;
	.loc 4 210 5
	min.s32 	%r83, %r81, %r71;
	.loc 4 238 5
	max.s32 	%r85, %r4, %r76;
	.loc 4 210 5
	min.s32 	%r86, %r85, %r39;
	.loc 3 42 1
	mad.lo.s32 	%r87, %r83, %r29, %r86;
	.loc 4 238 5
	max.s32 	%r89, %r8, %r76;
	.loc 4 210 5
	min.s32 	%r90, %r89, %r34;
	.loc 3 42 1
	mad.lo.s32 	%r91, %r87, %r30, %r90;
	cvta.to.global.u64 	%rd41, %rd7;
	.loc 3 42 1
	mul.wide.s32 	%rd42, %r91, 4;
	add.s64 	%rd43, %rd41, %rd42;
	ld.global.nc.f32 	%f54, [%rd43];

BB2_15:
	.loc 3 43 1
	sub.f32 	%f42, %f17, %f2;
	sub.f32 	%f43, %f18, %f2;
	mul.f32 	%f44, %f54, %f43;
	fma.rn.f32 	%f45, %f20, %f42, %f44;
	fma.rn.f32 	%f56, %f45, %f25, %f56;

BB2_16:
	.loc 3 46 1
	st.global.f32 	[%rd2], %f56;
	.loc 3 19 1
	add.s32 	%r97, %r97, %r10;
	add.s32 	%r96, %r96, %r10;
	add.s32 	%r95, %r95, %r10;
	add.s32 	%r94, %r94, %r10;
	add.s32 	%r93, %r93, %r10;
	setp.lt.s32 	%p13, %r98, %r28;
	@%p13 bra 	BB2_2;

BB2_17:
	.loc 3 48 2
	ret;
}


`
)
